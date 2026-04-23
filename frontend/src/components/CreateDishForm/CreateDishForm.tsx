import type { CreateDishFormProps } from './CreateDishForm.props';
import styles from './CreateDishForm.module.scss';
import { Input } from '@/elements/Input/Input';
import { Button } from '@/elements/Button/Button';
import { useState } from 'react';
import { createDish } from '@/api/Dish';

export const CreateDishForm = ({ categoryId, setOpen, fetchDishes }: CreateDishFormProps) => {
  const [errMsg, setErrMsg] = useState<string>('');
  const [successMsg, setSuccessMsg] = useState<string>('');
  const [disabled, setDisabled] = useState<boolean>(false);

  const [name, setName] = useState<string>('');
  const [desc, setDesc] = useState<string>('');
  const [recipe, setRecipe] = useState<string>('');
  const [timeToCook, setTimeToCook] = useState<number>(0);

  const create = async () => {
    try {
      setErrMsg('');
      setDisabled(true);
      await createDish({
        name: name,
        time_to_cook: timeToCook,
        recipe: recipe,
        description: desc,
        category_id: categoryId,
      });
      setName('');
      setDesc('');
      setRecipe('');
      setTimeToCook(0);
      setSuccessMsg('Блюдо успешно создано');
      fetchDishes(categoryId);
    } catch (err) {
      setErrMsg(err instanceof Error ? err.message : String(err));
      setSuccessMsg('');
    } finally {
      setDisabled(false);
    }
  };

  const setDishName = (val: unknown) => {
    setName(String(val));
  };

  const setDishDesc = (val: unknown) => {
    setDesc(String(val));
  };

  const setDishTimeToCook = (val: unknown) => {
    setTimeToCook(Number(val));
  };

  const setDishRecipe = (val: unknown) => {
    setRecipe(String(val));
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.bg} onClick={() => setOpen(false)}></div>
      <div className={styles.formWrapper}>
        <div className={styles.title}>Создание блюда</div>
        <div className={styles.inputWrapper}>
          <div className={styles.inputTitle}>Название блюда</div>
          <div className={styles.input}>
            <Input type="text" placeholder="Название блюда" value={name} onChange={setDishName} />
          </div>
        </div>
        <div className={styles.inputWrapper}>
          <div className={styles.inputTitle}>Рецепт</div>
          <div className={styles.input}>
            <Input type="textarea" placeholder="Рецепт" value={recipe} onChange={setDishRecipe} />
          </div>
        </div>
        <div className={styles.inputWrapper}>
          <div className={styles.inputTitle}>Время приготовления</div>
          <div className={styles.input}>
            <Input
              type="number"
              placeholder="Время приготовления"
              value={timeToCook}
              onChange={setDishTimeToCook}
            />
          </div>
        </div>
        <div className={styles.inputWrapper}>
          <div className={styles.inputTitle}>Описание</div>
          <div className={styles.input}>
            <Input type="textarea" placeholder="Описание" value={desc} onChange={setDishDesc} />
          </div>
        </div>
        {errMsg != '' && <div className={styles.errMsg}>{errMsg}</div>}
        {successMsg != '' && <div className={styles.successMsg}>{successMsg}</div>}
        <div className={styles.sendBtn}>
          <Button onClick={create} disabled={disabled}>
            Создать
          </Button>
        </div>
      </div>
    </div>
  );
};
