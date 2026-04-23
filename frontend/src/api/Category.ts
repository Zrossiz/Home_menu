import type { ICategory } from '@/types/Category';
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL;

export const getAllCategories = async (): Promise<ICategory[]> => {
  const { data } = await axios.get(`${API_URL}/category`);
  return data;
};

export const createCategory = async (name: string) => {
  await axios.post(`${API_URL}/category`, {
    name,
  });
};

export const updateCategory = async (name: string, categoryId: number) => {
  await axios.post(`${API_URL}/category/${categoryId}`, {
    data: {
      name,
    },
  });
};

export const deleteCategory = async (categoryId: number) => {
  await axios.delete(`${API_URL}/category/${categoryId}`);
};
