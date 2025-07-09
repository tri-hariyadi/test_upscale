import type { IApi } from 'lib/api/types';
import type { wrapPromise } from 'lib/utils.ts';

export interface TodoState {
  filterTodos: (tag: string, search: string) => Promise<void>;
  todosResources: ReturnType<typeof wrapPromise<Array<IApi['Todos']>>> | null;
  createTodoResource: () => void;
}
