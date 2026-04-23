import { Route, Routes } from 'react-router-dom';
import { Home } from '@/pages/Home/Home';
import { CategoryDishesPage } from '@/pages/CategoryDishes/CategoryDishes';
import { DishPage } from '@/pages/Dish/Dish';
import { Search } from './pages/Search/Search';

export const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/search" element={<Search />} />
      <Route path="/category/:categoryId" element={<CategoryDishesPage />} />
      <Route path="/category/:categoryId/dish/:dishId" element={<DishPage />} />
    </Routes>
  );
};
