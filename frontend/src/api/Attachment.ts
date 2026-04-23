import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL;

export const createAttachments = async (dishId: number, files: File[]) => {
  const formData = new FormData();

  files.forEach((file) => {
    formData.append('files', file);
  });

  await axios.post(`${API_URL}/attachment/dish/${dishId}`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};
