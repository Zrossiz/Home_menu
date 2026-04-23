import { useState } from 'react';
import styles from './CreateCategoryForm.module.scss';
import type { CreateCategoryFormProps } from './CreateCategoryForm.props';
import { createCategory } from '@/api/Category';
import { Input } from '@/elements/Input/Input';
import { Button } from '@/elements/Button/Button';

export const CreateCategoryForm = ({ setOpen, fetchCategories }: CreateCategoryFormProps) => {
  const [name, setName] = useState<string>('');
  const [disabled, setDisabled] = useState<boolean>(false);
  const [errMsg, setErrMsg] = useState<string>('');
  const [successMsg, setSuccessMsg] = useState<string>('');

  const create = async () => {
    try {
      setDisabled(true);
      if (name == '') {
        throw new Error('имя не может быть пустым');
      }
      setErrMsg('');
      setSuccessMsg('');
      await createCategory(name);
      setName('');
      setSuccessMsg('Категория успешно создана');
      fetchCategories();
    } catch (err) {
      console.log(err);
      setErrMsg(err instanceof Error ? err.message : String(err));
      setSuccessMsg('');
    } finally {
      setDisabled(false);
    }
  };

  const setInputCategoryName = (val: unknown) => {
    setName(String(val));
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.bg} onClick={() => setOpen(false)}></div>
      <div className={styles.formWrapper}>
        <div className={styles.title}>Создание категории</div>
        <div className={styles.inputWrapper}>
          <div className={styles.inputTitle}>Название категории</div>
          <div className={styles.input}>
            <Input
              type="text"
              placeholder="Название категории"
              value={name}
              onChange={setInputCategoryName}
            />
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
