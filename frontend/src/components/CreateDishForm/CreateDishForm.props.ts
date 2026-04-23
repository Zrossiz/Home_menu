export interface CreateDishFormProps {
  categoryId: number;
  setOpen(arg0: boolean): void;
  fetchDishes(id: number): Promise<void>;
}
