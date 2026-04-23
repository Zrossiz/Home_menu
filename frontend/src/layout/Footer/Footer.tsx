import styles from '@/layout/Footer/Footer.module.scss';

export const Footer = () => {
  return (
    <footer className={styles.footer}>
      <div className={styles.wrapper}>
        <span className={styles.title}>От Шнапика с любовью</span>
      </div>
    </footer>
  );
};
