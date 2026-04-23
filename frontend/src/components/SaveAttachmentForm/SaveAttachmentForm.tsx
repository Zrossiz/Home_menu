import { useState } from 'react';
import { createAttachments } from '@/api/Attachment';
import { Button } from '@/elements/Button/Button';
import type { SaveAttachmentFormProps } from './SaveAttachmentForm.props';
import styles from './SaveAttachmentForm.module.scss';

export const SaveAttachmentForm = ({ dishId, fetchDish }: SaveAttachmentFormProps) => {
  const [attachments, setAttachments] = useState<File[]>([]);

  const saveAttachments = async () => {
    try {
      await createAttachments(dishId, attachments);
      fetchDish(dishId);
      setAttachments([]);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.title}>Загрузить фото</div>

      <input
        type="file"
        multiple
        onChange={(e) => {
          if (!e.target.files) return;
          const selectedFiles = Array.from(e.target.files);
          setAttachments((prev) => [...prev, ...selectedFiles]);
        }}
      />
      <div className={styles.preview}>
        {attachments.map((file, i) => (
          <img key={i} src={URL.createObjectURL(file)} />
        ))}
      </div>

      <Button onClick={saveAttachments}>Сохранить</Button>
    </div>
  );
};
