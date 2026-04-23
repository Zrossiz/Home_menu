import { DishListItem } from '@/components/DishListItem/DishListItem';
import styles from '@/pageComponents/DishesList/DishesList.module.scss';
import type { IDish } from '@/types/Dish';
import { CreateDishListItem } from '@/components/CreateDishListItem/CreateDishListItem';

interface IDishesListProps {
  dishes: IDish[];
  categoryId: number;
  fetchDishes(id: number): Promise<void>;
}

export const DishesList = ({ dishes, categoryId, fetchDishes }: IDishesListProps) => {
  return (
    <div className={styles.listWrapper}>
      <ul>
        <CreateDishListItem fetchDishes={fetchDishes} categoryId={categoryId} />
        {dishes.map((item) => {
          return <DishListItem key={item.id} item={item} categoryId={Number(categoryId)} />;
        })}
      </ul>
    </div>
  );
};
