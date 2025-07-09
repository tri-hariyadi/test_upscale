import { forwardRef, useImperativeHandle, useRef, useState } from 'react';

import { Modal } from 'components';

import TodoFormComponent from './TodoFormComponent.tsx';
import type { TodoFormProps, TodoFormRefObject } from './types';

import type { ModalRefObject } from '../../../components/Modal/types';

const TodoForm = forwardRef<TodoFormRefObject>((_, ref) => {
  const modalRef = useRef<ModalRefObject>(null);
  const [isEdited, setIsEdited] = useState(false);
  const [formValue, setFormValue] = useState<TodoFormProps['formValue']>(undefined);

  useImperativeHandle(
    ref,
    () => ({
      toggle(open, isEdit, formValue) {
        modalRef.current?.toggle(open);
        setIsEdited(isEdit || false);
        setFormValue(formValue);
      }
    }),
    []
  );

  return (
    <Modal
      ref={modalRef}
      title={isEdited ? 'Edit Todo' : 'Create Todo'}
      className="w-full md:max-w-2xl px-4 sm:px-10 py-8 sm:py-10"
    >
      <TodoFormComponent
        isEdited={isEdited}
        formValue={formValue}
        onClose={(value) => modalRef.current?.toggle(value)}
      />
    </Modal>
  );
});

export default TodoForm;
