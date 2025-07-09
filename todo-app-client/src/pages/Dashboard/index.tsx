import { Suspense } from 'react';

import { CirclePlus } from 'lucide-react';

import { Button, Fallback, Profile } from 'components';
import type { IApi } from 'lib/api/types';

import Layout, { LayoutLeft, LayoutRight } from './layout.tsx';
import ListTask from './ListTask';
import Quote from './Quote';
import TodoForm from './TodoForm';
import useData from './useData';

const Dashboard = () => {
  const { todoForm, todayDate } = useData();

  const onEdit = (value: IApi['Todos']) => {
    todoForm.current?.toggle(true, true, {
      id: value.id,
      title: value.title,
      description: value.description,
      tag: value.tag,
      due_date: new Date(value.due_date)
    });
  };

  return (
    <>
      <Layout>
        <LayoutLeft className="shadow-lg px-4 py-6 rounded-xl border-1 border-neutral-200">
          <Profile />
        </LayoutLeft>

        <LayoutRight>
          <div className="shadow-lg px-4 py-6 rounded-xl border-1 border-neutral-200 flex flex-col mb-8 bg-white">
            <span className="text-xl font-inter-bold mb-4">Today's quotes {todayDate()}</span>
            <Suspense fallback={<Fallback />}>
              <Quote />
            </Suspense>
            <Button
              className="self-end flex gap-2 justify-center items-center"
              onClick={() => todoForm.current?.toggle(true)}
            >
              <CirclePlus className="h-5 w-5" />
              Create Todo
            </Button>
          </div>
          <Suspense fallback={<Fallback />}>
            <ListTask onEdit={onEdit} />
          </Suspense>
        </LayoutRight>
      </Layout>

      <TodoForm ref={todoForm} />
    </>
  );
};

export default Dashboard;
