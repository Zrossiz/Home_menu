export interface IDish {
  id: number;
  name: string;
  time_to_cook: number;
  recipe: string;
  description: string;
  category_id: number;
  created_at: Date;
}

export interface IDishWithImages extends IDish {
  images: string[];
}

export interface ICreateDish {
  name: string;
  time_to_cook: number;
  recipe: string;
  description: string;
  category_id: number;
}
