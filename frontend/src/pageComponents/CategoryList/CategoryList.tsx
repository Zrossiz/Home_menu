import { CategoryListItem } from '@/components/CategoryListItem/CategoryListItem';
import { CreateCategoryItem } from '@/components/CreateCategoryItem/CreateCategoryItem';
import styles from '@/pageComponents/CategoryList/CategoryList.module.scss';
import type { ICategory } from '@/types/Category';

interface ICategoryListProps {
  categories: ICategory[];
  fetchCategories: () => Promise<void>;
}

export const CategoryList = ({ categories, fetchCategories }: ICategoryListProps) => {
  return (
    <div className={styles.listWrapper}>
      <ul>
        <CreateCategoryItem fetchCategories={fetchCategories} />
        {categories.map((item) => {
          return <CategoryListItem key={item.id} id={item.id} name={item.name} />;
        })}
      </ul>
    </div>
  );
};
