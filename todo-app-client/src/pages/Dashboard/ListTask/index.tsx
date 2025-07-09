import { OctagonX } from 'lucide-react';

import { Show, TodoCard } from 'components';
import type { IApi } from 'lib/api/types';
import { useTodosStore } from 'store/TodosStore';

const ListTask = ({ onEdit }: { onEdit: (value: IApi['Todos']) => void }) => {
  const todosResources = useTodosStore((s) => s.todosResources);
  const todos = todosResources?.read() || [];

  return (
    <div>
      <Show>
        <Show.When isTrue={todos.length}>
          {todos?.map((todo, idx) => (
            <TodoCard
              key={`todos-${idx}`}
              todo={todo}
              onEdit={() => onEdit(todo)}
              className={todos?.length - 1 !== idx ? 'mb-8' : ''}
            />
          ))}
        </Show.When>
        <Show.Else>
          <span className="text-neutral-500 flex flex-col gap-3 justify-center items-center text-lg">
            <OctagonX className="w-10 h-10" />
            There is no todos data yet
          </span>
        </Show.Else>
      </Show>
    </div>
  );
};

export default ListTask;
