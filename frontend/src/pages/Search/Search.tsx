import { Layout } from '@/layout/Layout/Layout';
import type { IDish } from '@/types/Dish';
import { useState } from 'react';
import styles from './Search.module.scss';
import { Input } from '@/elements/Input/Input';
import { findDishes } from '@/api/Dish';
import { DishListItem } from '@/components/DishListItem/DishListItem';

export const Search = () => {
  const [dishes, setDishes] = useState<IDish[]>([]);
  const [pattern, setPattern] = useState<string>('');

  const setSearchPattern = async (pat: unknown) => {
    try {
      setPattern(String(pat));

      const res = await findDishes(pattern);
      if (res.status == 200 && res.data.length > 0) {
        setDishes(res.data);
      }
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Layout title={'Домашнее меню | Поиск'}>
      <div className={styles.wrapper}>
        <div className={styles.inputWrapper}>
          <Input
            placeholder={'Название блюда'}
            type={'text'}
            value={pattern}
            onChange={setSearchPattern}
          />
        </div>
      </div>
      <div className={styles.dishesWrapper}>
        {dishes.length > 0 ? (
          <div className={styles.listWrapper}>
            {dishes.map((dish) => {
              return <DishListItem key={dish.id} item={dish} categoryId={dish.category_id} />;
            })}
          </div>
        ) : (
          <div className={styles.notFound}>Ничего не найдено 🥲</div>
        )}
      </div>
    </Layout>
  );
};
