import styles from '@/components/CategoryListItem/CategoryListItem.module.scss';

interface CategoryListItemProps {
  id: number;
  name: string;
}

export const CategoryListItem = ({ id, name }: CategoryListItemProps) => {
  return (
    <li className={styles.li}>
      <a className={styles.link} href={`/category/${id}`}>
        {name}
      </a>
    </li>
  );
};
