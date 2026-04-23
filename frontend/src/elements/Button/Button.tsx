import type { ButtonProps } from './Button.props';
import styles from './Button.module.scss';

export const Button = ({ children, onClick = () => {}, disabled }: ButtonProps) => {
  return (
    <button className={styles.btn} disabled={disabled} onClick={(arg0) => onClick(arg0)}>
      {children}
    </button>
  );
};
