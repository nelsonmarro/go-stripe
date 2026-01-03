import {
  loadStripe,
  StripeCardElement,
  StripeElements,
  StripeElementStyle,
} from "@stripe/stripe-js";
import { VTState } from "./models";

// Make the card element globally accessible
declare global {
  interface Window {
    stripeCardElement: StripeCardElement | null;
    processPayment: (state: VTState) => Promise<void>;
    initStripeCard: () => Promise<void>;
  }
}

function showCardError(msg: string) {
  const cardMessages = document.getElementById(
    "card-messages",
  ) as HTMLDivElement;
  if (cardMessages) {
    cardMessages.classList.remove(
      "hidden",
      "text-green-700",
      "bg-green-100",
      "border-green-400",
    );
    cardMessages.classList.add("text-red-700", "bg-red-100", "border-red-400");
    cardMessages.innerText = msg;
  }
}

function showCardSuccess() {
  const cardMessages = document.getElementById(
    "card-messages",
  ) as HTMLDivElement;
  if (cardMessages) {
    cardMessages.classList.remove(
      "hidden",
      "text-red-700",
      "bg-red-100",
      "border-red-400",
    );
    cardMessages.classList.add(
      "text-green-700",
      "bg-green-100",
      "border-green-400",
    );
    cardMessages.innerText = "Payment successful! Thank you.";
  }
}

window.processPayment = async (state: VTState) => {
  console.log("Processing payment with state:", state);

  const stripe = await stripePromise;
  const cardElement = window.stripeCardElement;

  if (!stripe || !cardElement) {
    console.error("Stripe or Card Element not initialized.");
    return;
  }

  if (!state.payment_intent || !state.payment_intent.client_secret) {
    console.error("No payment intent client secret available.");
    return;
  }

  try {
    const result = await stripe.confirmCardPayment(
      state.payment_intent.client_secret,
      {
        payment_method: {
          card: cardElement,
          billing_details: {
            name: state.form.cardHolder,
            email: state.form.email,
          },
        },
      },
    );

    if (result.error) {
      // card was declined or other error
      showCardError(result.error?.message ?? "");
    } else if (result.paymentIntent) {
      if (result.paymentIntent.status == "succeeded") {
        state.payment_method =
          result.paymentIntent.payment_method_types[0] ?? "";
        state.payment_amount = result.paymentIntent.amount.toString();
        state.payment_currency = result.paymentIntent.currency;

        showCardSuccess();

        // Dispatch success event
        const form = document.getElementById("payment-form");
        if (form) {
          form.dispatchEvent(
            new CustomEvent("payment-success", {
              detail: { paymentIntent: result.paymentIntent },
              bubbles: true,
            }),
          );
        }
      }
    }
  } catch (err) {
    console.error("Error processing payment:", err);
    showCardError("Invalid payment attempt. Please try again.");
  } finally {
    // Reset signals to update UI reactively via Data-Star
    state.is_processing = false;
    state.payment_intent_success = false;
    state.payment_intent = null;
  }
};

// --- Stripe Initialization ---
const getStripeKey = (): string => {
  const meta = document.querySelector('meta[name="stripe-key"]');
  return meta ? meta.getAttribute("content") || "" : "";
};

const stripePromise = loadStripe(getStripeKey());

window.initStripeCard = async () => {
  const stripe = await stripePromise;
  if (!stripe) {
    console.error("Failed to load Stripe.js");
    return;
  }

  // Ensure target element exists
  if (!document.getElementById("card-element")) {
    console.warn("Card element container not found.");
    return;
  }

  // Clean up existing card if any
  if (window.stripeCardElement) {
    window.stripeCardElement.unmount();
    window.stripeCardElement.destroy();
    window.stripeCardElement = null;
  }

  const elements: StripeElements = stripe.elements();
  const style: StripeElementStyle = {
    base: {
      fontSize: "16px",
      lineHeight: "24px",
    },
  };

  const card = elements.create("card", {
    style: style,
    hidePostalCode: true,
  });

  card.mount("#card-element");
  window.stripeCardElement = card;

  // check for input errors
  card.on("change", function (event) {
    const displayError = document.getElementById("card-errors");
    if (displayError) {
      if (event.error) {
        displayError.classList.remove("hidden");
        displayError.textContent = event.error.message;
      } else {
        displayError.classList.add("hidden");
        displayError.textContent = "";
      }
    }
  });
  console.log("Stripe Card Element initialized");
};

// Inicialize Stripe Card Element on DOM load
document.addEventListener("DOMContentLoaded", () => {
  if (document.getElementById("card-element")) {
    window.initStripeCard();
  }
});
