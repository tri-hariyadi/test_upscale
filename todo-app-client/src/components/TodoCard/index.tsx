import { SquarePen, Trash2 } from 'lucide-react';
import { toast } from 'react-toastify';
import Swal from 'sweetalert2';

import api from 'lib/api/api.ts';
import type { IApi } from 'lib/api/types';
import { cn, formatDate } from 'lib/utils.ts';
import { useTodosStore } from 'store/TodosStore';

import { Button } from '../index.ts';
import Show from '../Show';

interface TodoCardProps {
  onEdit?: () => void;
  todo: IApi['Todos'];
  className?: string;
}

const TodoCard = ({ onEdit, todo, className }: TodoCardProps) => {
  const createTodoResource = useTodosStore((state) => state.createTodoResource);

  const deleteTodo = async () => {
    const resp = await api.deleteTodo(todo.id);
    if (resp.status) {
      toast.success('Delete todo has been successful');
      createTodoResource();
    } else {
      toast.success(resp.message || 'Delete todo failed');
    }
  };

  const onDelete = () =>
    Swal.fire({
      title: 'Are you sure you want to delete todo?',
      showCancelButton: true,
      confirmButtonText: 'Yes',
      confirmButtonColor: 'oklch(57.7% 0.245 27.325)'
    }).then(async (result) => {
      if (result.isConfirmed) {
        deleteTodo();
      }
    });

  return (
    <div className={cn('shadow-lg px-4 py-6 rounded-xl border-1 border-neutral-200 flex flex-col bg-white', className)}>
      <div className="flex justify-between mb-2">
        <h2 className="text-xl font-inter-bold">{todo.title}</h2>
        <Show>
          <Show.When isTrue={todo.tag === 'active'}>
            <span className="bg-red-100 text-red-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded-sm self-baseline">
              Not Finished
            </span>
          </Show.When>
          <Show.Else>
            <span className="bg-green-100 text-green-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded-sm self-baseline">
              Finished
            </span>
          </Show.Else>
        </Show>
      </div>
      <p className="text-sm mb-2 whitespace-pre-wrap">{todo.description}</p>
      <span className="text-sm inline-block mb-1">
        <strong className="text-red-600 font-inter-semibold">Due Date:</strong> {formatDate(todo.due_date)}
      </span>
      <span className="text-sm">
        <strong className="font-inter-semibold">Updated At:</strong> {formatDate(todo.updated_at)}
      </span>
      <div className="self-end flex gap-4 mt-6">
        <Button variant="secondary" className="self-baseline flex gap-2 items-center" onClick={onEdit}>
          <SquarePen className="h-5 w-5" />
          Edit
        </Button>
        <Button variant="danger" className="self-baseline flex gap-2 justify-center items-center" onClick={onDelete}>
          <Trash2 className="h-5 w-5" />
          Delete
        </Button>
      </div>
    </div>
  );
};

export default TodoCard;
