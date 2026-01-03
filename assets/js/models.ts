import { PaymentIntent } from "@stripe/stripe-js";

export interface VTForm {
  amount: string;
  cardHolder: string;
  email: string;
}

export interface VTState {
  form: VTForm;
  errors: VTForm;
  form_submitted: boolean;
  is_processing: boolean;
  payment_intent_success: boolean;
  payment_intent: PaymentIntent | null;
  payment_method: string;
  payment_amount: string;
  payment_currency: string;
}
