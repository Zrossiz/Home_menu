import type { ICreateDish, IDish, IDishWithImages } from '@/types/Dish';
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL;

export const getDishesByCategory = async (id: number): Promise<IDish[]> => {
  const { data } = await axios.get(`${API_URL}/dish/category/${id}`);

  return data;
};

export const getDish = async (id: number): Promise<IDishWithImages> => {
  const { data } = await axios.get(`${API_URL}/dish/${id}`);
  return data;
};

export const createDish = async (data: ICreateDish) => {
  await axios.post(`${API_URL}/dish/`, data);
};

export const findDishes = async (pattern: string) => {
  return axios.get(`${API_URL}/dish/search?search=${pattern}`);
};
