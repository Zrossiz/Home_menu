import { getDishesByCategory } from '@/api/Dish';
import { Layout } from '@/layout/Layout/Layout';
import { DishesList } from '@/pageComponents/DishesList/DishesList';
import styles from '@/pages/CategoryDishes/CategoryDishes.module.scss';
import type { IDish } from '@/types/Dish';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

export const CategoryDishesPage = () => {
  const { categoryId } = useParams();
  const [dishes, setDishes] = useState<IDish[]>([]);

  const fetchDishes = async (id: number) => {
    try {
      (async () => {
        const data = await getDishesByCategory(id);
        if (data && data.length > 0) {
          setDishes(data);
        }
      })();
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    fetchDishes(Number(categoryId));
  }, []);

  return (
    <Layout title="Домашнее меню | Блюда">
      <div className={styles.wrapper}>
        <DishesList fetchDishes={fetchDishes} dishes={dishes} categoryId={Number(categoryId)} />
      </div>
    </Layout>
  );
};
