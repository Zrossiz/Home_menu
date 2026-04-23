export interface InputProps {
  placeholder?: string;
  type: 'text' | 'phone' | 'password' | 'number' | 'textarea';
  value?: unknown;
  onChange?: (arg0: unknown) => void;
  defaultValue?: unknown;
}
