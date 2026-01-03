// assets/js/validators.ts
import { VTForm } from "./models";

export function validatePaymentForm(form: VTForm): VTForm {
  const errors: VTForm = { amount: "", cardHolder: "", email: "" };

  // Amount validation
  if (!form.amount) {
    errors.amount = "Amount is required.";
  } else if (isNaN(Number(form.amount)) || Number(form.amount) <= 0) {
    errors.amount = "Please enter a valid, positive number.";
  }

  // Card Holder validation
  if (!form.cardHolder.trim()) {
    errors.cardHolder = "Card holder name is required.";
  }

  // Email validation
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!form.email) {
    errors.email = "Email is required.";
  } else if (!emailRegex.test(form.email)) {
    errors.email = "Please enter a valid email address.";
  }

  return errors;
}

// Expose the function to the global window object so Data-Star can call it
declare global {
  interface Window {
    validatePaymentForm: (form: VTForm) => VTForm;
  }
}
window.validatePaymentForm = validatePaymentForm;
