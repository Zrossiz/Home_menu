export interface SaveAttachmentFormProps {
  dishId: number;
  fetchDish: (dishId: number) => Promise<void>;
}
