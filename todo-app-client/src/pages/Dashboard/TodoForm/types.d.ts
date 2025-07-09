export interface TodoFormProps {
  isEdited: boolean;
  onClose: (value: boolean) => void;
  formValue?: {
    id: number;
    title: string;
    description: string;
    tag: string;
    due_date: Date;
  };
}

export interface TodoFormRefObject {
  toggle: (open: boolean, isEdit?: boolean, formValue?: TodoFormProps['formValue']) => void;
}
