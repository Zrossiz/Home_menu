import { Header } from '@/layout/Header/Header';
import { Footer } from '@/layout/Footer/Footer';
import styles from './Layout.module.scss';
import { useEffect } from 'react';

interface LayoutProps {
  children: React.ReactNode;
  title: string;
}

export const Layout = ({ children, title }: LayoutProps) => {
  useEffect(() => {
    document.title = title;
  }, []);

  return (
    <div className={styles.wrapper}>
      <Header />
      <main>{children}</main>
      <Footer />
    </div>
  );
};
