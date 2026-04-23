import { getAllCategories } from '@/api/Category';
import { Layout } from '@/layout/Layout/Layout';
import { CategoryList } from '@/pageComponents/CategoryList/CategoryList';
import styles from '@/pages/Home/Home.module.scss';
import type { ICategory } from '@/types/Category';
import { useEffect, useState } from 'react';

export const Home = () => {
  const [categories, setCategories] = useState<ICategory[]>();

  const fetchCategories = async () => {
    try {
      const categoriesApi = await getAllCategories();
      setCategories(categoriesApi);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    fetchCategories();
  }, []);

  return (
    <Layout title="Домашнее меню | Категории">
      <div className={styles.wrapper}>
        <CategoryList categories={categories ?? []} fetchCategories={fetchCategories} />
      </div>
    </Layout>
  );
};
