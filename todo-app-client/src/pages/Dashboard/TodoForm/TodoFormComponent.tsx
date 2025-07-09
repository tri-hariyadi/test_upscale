import { type ChangeEvent } from 'react';

import { Save, X } from 'lucide-react';
import { toast } from 'react-toastify';

import { Button, InputText, RadioButton } from 'components';
import api from 'lib/api/api.ts';
import type { IApi } from 'lib/api/types';
import { useForm } from 'lib/hooks/useForm.ts';
import { schema } from 'lib/validator.ts';
import { useSearchStore } from 'store/SearchStore';
import { useTodosStore } from 'store/TodosStore';

import type { TodoFormProps } from './types';

import InputDateTime from '../../../components/InputDateTime';

const TodoFormComponent = ({
  isEdited,
  onClose,
  formValue = {
    id: 0,
    title: '',
    description: '',
    tag: '',
    due_date: new Date()
  }
}: TodoFormProps) => {
  const createTodoResource = useTodosStore((state) => state.createTodoResource);
  const setTag = useSearchStore((state) => state.setTag);
  const { values, errors, handleSubmit, onChange, isLoading, setIsLoading } = useForm(formValue, {
    title: (e) => schema.string().label('email').min(5).validate(e),
    description: (e) => schema.string().label('description').min(5).validate(e),
    tag: (e) => schema.string().label('task status').validate(e)
  });

  const createTodo = async (payload: IApi['TodosRequest']) => {
    setIsLoading(true);
    const resp = await api.createTodo(payload);
    if (resp.status) {
      setTag('');
      toast.success('Add todo successfully');
      createTodoResource();
      onClose(false);
    }
    setIsLoading(false);
  };

  const updateTodo = async (payload: IApi['TodosRequest']) => {
    setIsLoading(true);
    const resp = await api.updateTodo(formValue.id, payload);
    if (resp.status) {
      setTag('');
      toast.success('Update todo successfully');
      createTodoResource();
      onClose(false);
    }
    setIsLoading(false);
  };

  const onSubmit = () =>
    handleSubmit((value) => {
      if (!value) return;
      if (isEdited) {
        return updateTodo(value);
      }
      createTodo(value);
    });

  return (
    <>
      <InputText
        type="text"
        placeholder="Enter todo title"
        name="title"
        id="title"
        label="Title"
        value={values.title}
        onChange={(e: ChangeEvent<HTMLInputElement>) => onChange({ title: e.target.value })}
        error={errors.title}
      />
      <InputDateTime
        value={values.due_date}
        id="duedate"
        error={errors.due_date}
        label="Due Date"
        onChange={(e) => onChange({ due_date: e })}
      />
      <RadioButton
        data={[
          { label: 'Finished', value: 'inactive' },
          { label: 'Not Finished', value: 'active' }
        ]}
        name="tag"
        label="Task Status"
        error={errors.tag}
        value={values.tag}
        onChange={(e) => onChange({ tag: e })}
      />
      <InputText
        type="textarea"
        placeholder="Enter todo description"
        name="description"
        id="description"
        label="Description"
        value={values.description}
        onChange={(e: ChangeEvent<HTMLTextAreaElement>) => onChange({ description: e.target.value })}
        error={errors.description}
      />
      <div className="flex gap-2 items-center justify-end mt-6">
        <Button
          variant="danger"
          className="self-end flex gap-2 justify-center items-center"
          onClick={() => onClose(false)}
        >
          <X className="h-5 w-5" />
          Cancel
        </Button>
        <Button className="self-end flex gap-2 justify-center items-center" isLoading={isLoading} onClick={onSubmit}>
          <Save className="h-5 w-5" />
          Submit
        </Button>
      </div>
    </>
  );
};

export default TodoFormComponent;
