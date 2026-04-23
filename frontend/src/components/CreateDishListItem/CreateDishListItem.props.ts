export interface CreateDishListItemProps {
  categoryId: number;
  fetchDishes(id: number): Promise<void>;
}
