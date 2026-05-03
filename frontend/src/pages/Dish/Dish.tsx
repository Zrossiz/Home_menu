import { getDish } from '@/api/Dish';
import { DishInfo } from '@/components/DishInfo/DishInfo';
import { Gallery } from '@/components/Gallery/Gallery';
import { SaveAttachmentForm } from '@/components/SaveAttachmentForm/SaveAttachmentForm';
import { Layout } from '@/layout/Layout/Layout';
import styles from '@/pages/Dish/Dish.module.scss';
import type { IDishWithImages } from '@/types/Dish';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

export const DishPage = () => {
  const { dishId } = useParams();
  const [dish, setDish] = useState<IDishWithImages>();

  const fetchDish = async (dishId: number) => {
    try {
      const dishApi = await getDish(dishId);
      console.log(dishApi)
      setDish(dishApi);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    fetchDish(Number(dishId));
  }, []);

  return (
    <Layout title="Домашнее меню | Блюдо">
      {dish && (
        <div className={styles.wrapper}>
          <Gallery images={dish?.images ?? []} />
          <DishInfo item={dish || null} />
          <SaveAttachmentForm fetchDish={fetchDish} dishId={dish.id} />
        </div>
      )}
    </Layout>
  );
};
