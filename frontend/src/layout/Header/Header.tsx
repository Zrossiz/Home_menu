import styles from '@/layout/Header/Header.module.scss';

export const Header = () => {
  return (
    <header className={styles.header}>
      <div className={styles.wrapper}>
        <a href="/" className={styles.title}>
          Домашнее меню
        </a>
        <div className={styles.searchWrapper}>
          <a href="/search">
            <img src="/loop.svg" alt="Поиск" />
          </a>
        </div>
      </div>
    </header>
  );
};
