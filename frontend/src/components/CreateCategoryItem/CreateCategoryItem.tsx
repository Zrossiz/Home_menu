import styles from './CreateCategoryItem.module.scss';
import { useState } from 'react';
import { CreateCategoryForm } from '../CreateCategoryForm/CreateCategoryForm';
import type { CreateCategoryItemProps } from './CreateCategoryItem.props';

export const CreateCategoryItem = ({ fetchCategories }: CreateCategoryItemProps) => {
  const [open, setOpen] = useState<boolean>(false);

  return (
    <li className={styles.li}>
      {open && <CreateCategoryForm fetchCategories={fetchCategories} setOpen={setOpen} />}
      <div className={styles.createItem} onClick={() => setOpen(true)}>
        Создать
      </div>
    </li>
  );
};
