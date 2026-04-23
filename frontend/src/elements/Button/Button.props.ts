import type { ReactNode } from 'react';

export interface ButtonProps {
  onClick?: (arg: unknown) => void;
  children: ReactNode;
  disabled?: boolean;
}
