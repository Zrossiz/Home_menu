import type { IDishWithImages } from '@/types/Dish';
import styles from '@/components/DishInfo/DishInfo.module.scss';

interface IDishInfoProps {
  item: IDishWithImages | null;
}

export const DishInfo = ({ item }: IDishInfoProps) => {
  console.log(item);
  return (
    <>
      <div className={styles.wrapper}>
        <div className={styles.title}>{item?.name}</div>

        <div className={styles.row}>
          Время приготовления: <span>{item?.time_to_cook} мин</span>
        </div>

        <div className={styles.description}>{item?.description}</div>

        <div className={styles.recipe}>{item?.recipe}</div>
      </div>
    </>
  );
};
