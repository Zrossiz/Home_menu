import type { IDish } from '@/types/Dish';
import styles from '@/components/DishListItem/DishListItem.module.scss';

interface DishesListProps {
  item: IDish;
  categoryId?: number;
}

export const DishListItem = ({ item, categoryId }: DishesListProps) => {
  return (
    <li className={styles.li}>
      <a className={styles.link} href={`/category/${categoryId}/dish/${item.id}`}>
        <div className={styles.imgWrapper}>
          <img 
            src={`${import.meta.env.VITE_FILE_PATH}${item.image}`} 
            alt={item.name} 
          />
        </div>
        {item.name}
      </a>
    </li>
  );
};
