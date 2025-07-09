import { useEffect, useRef } from 'react';

import { useTodosStore } from 'store/TodosStore';

import type { TodoFormRefObject } from './TodoForm/types';

const useData = () => {
  const createTodoResource = useTodosStore((state) => state.createTodoResource);
  const todoForm = useRef<TodoFormRefObject>(null);

  const todayDate = () => {
    const today = new Date();
    return today.toLocaleDateString('en-GB', {
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    });
  };

  useEffect(() => {
    createTodoResource();
  }, []);

  return {
    createTodoResource,
    todoForm,
    todayDate
  };
};

export default useData;
