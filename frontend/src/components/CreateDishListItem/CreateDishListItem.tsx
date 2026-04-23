import { useState } from 'react';
import styles from './CreateDishListItem.module.scss';
import type { CreateDishListItemProps } from './CreateDishListItem.props';
import { CreateDishForm } from '../CreateDishForm/CreateDishForm';

export const CreateDishListItem = ({ categoryId, fetchDishes }: CreateDishListItemProps) => {
  const [open, setOpen] = useState<boolean>(false);

  return (
    <li className={styles.li}>
      {open && (
        <CreateDishForm fetchDishes={fetchDishes} setOpen={setOpen} categoryId={categoryId} />
      )}
      <div className={styles.createItem} onClick={() => setOpen(true)}>
        Создать
      </div>
    </li>
  );
};
