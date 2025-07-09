import { create } from 'zustand/react';

import api from 'lib/api/api.ts';
import { wrapPromise } from 'lib/utils.ts';
import type { TodoState } from 'store/TodosStore/types';

export const useTodosStore = create<TodoState>((set) => {
  return {
    todosResources: null,
    createTodoResource: () => {
      const promise = api.getAllTodo().then((resp) => {
        return resp.data ?? [];
      });
      const wrapped = wrapPromise(promise);
      set({ todosResources: wrapped });
    },
    // getTodosResource: () => {
    //   if (!todoResource) {
    //     const promise = api.getAllTodo().then((resp) => {
    //       if (resp.status && resp.data?.length) {
    //         set({ todos: resp.data });
    //       }
    //       return resp.data ?? [];
    //     });
    //     todoResource = wrapPromise(promise);
    //   }
    //   return todoResource;
    // },
    filterTodos: async (tag, search) => {
      const promise = api.filterTodo(search, tag).then((resp) => {
        return resp.data ?? [];
      });
      const wrapped = wrapPromise(promise);
      set({ todosResources: wrapped });
    }
  };
});
