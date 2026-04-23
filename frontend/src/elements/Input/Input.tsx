import type { InputProps } from './Input.props';
import styles from './Input.module.scss';

export const Input = ({
  placeholder,
  value,
  type,
  defaultValue,
  onChange = () => {},
}: InputProps) => {
  switch (type) {
    case 'textarea':
      return (
        <textarea
          className={styles.input}
          value={value as string}
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          defaultValue={defaultValue as string}
        />
      );
    case 'text':
      return (
        <input
          className={styles.input}
          value={value as string}
          type={'text'}
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          defaultValue={defaultValue as string}
        />
      );
    case 'password':
      return (
        <input
          className={styles.input}
          value={value as string}
          type={'password'}
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          defaultValue={defaultValue as string}
        />
      );
    case 'number':
      return (
        <input
          className={styles.input}
          value={value as number}
          type={'number'}
          placeholder={placeholder}
          onChange={(e) => onChange(+e.target.value)}
          defaultValue={defaultValue as number}
        />
      );
  }
};
