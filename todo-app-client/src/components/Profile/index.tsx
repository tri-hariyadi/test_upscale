import { type ChangeEvent, useEffect, useRef } from 'react';

import { LayoutList, List, ListChecks, LogOut, Search } from 'lucide-react';

import DummyProfile from 'assets/dummy_profile.png';
import { Button } from 'components';
import api from 'lib/api/api.ts';
import { cn } from 'lib/utils.ts';
import { useUserStore } from 'store/AuthStore';
import { useSearchStore } from 'store/SearchStore';
import { useTodosStore } from 'store/TodosStore';

const Profile = () => {
  const user = useUserStore((state) => state.user);
  const setUser = useUserStore((state) => state.setUser);
  const filterTodos = useTodosStore((state) => state.filterTodos);
  const search = useSearchStore();
  const isMounted = useRef(false);
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);

  const handleLogout = async () => {
    api.logout();
    window.localStorage.removeItem('access_token');
    window.location.replace('/sign-in');
    setUser(null);
  };

  useEffect(() => {
    if (isMounted.current) {
      filterTodos(search.search.tag, search.search.keyword);
    }
    isMounted.current = true;
  }, [search.search.tag]);

  useEffect(() => {
    return () => {
      search.setKeyword('');
      search.setTag('');
    };
  }, []);

  const onSearch = (e: ChangeEvent<HTMLInputElement>) => {
    search.setKeyword(e.target.value);
    if (timeoutRef.current) clearTimeout(timeoutRef.current);

    timeoutRef.current = setTimeout(() => {
      filterTodos(search.search.tag, search.search.keyword);
    }, 600);
  };

  const btnClassName = 'text-black flex gap-2 items-center px-4 py-2 rounded-md hover:bg-[#EDEBEB]';

  return (
    <div className="flex flex-col bg-white min-w-72 md:min-w-full">
      <img
        src={DummyProfile}
        alt="foto-profile"
        className="object-fill w-[120px] h-[120px] relative left-1/2 -translate-x-1/2 mb-4"
      />
      <h2 className="text-center font-inter-bold text-[18px] capitalize">Hay, {user?.name}</h2>
      <h3 className="text-center text-sm mt-1">{user?.email}</h3>
      <h2 className="mt-8 mb-4 font-inter-semibold">TODOS</h2>
      <div className="hidden bg-[#F6F5F5] rounded-lg md:flex justify-center items-center px-4 py-2 outline outline-1 outline-transparent focus-within:outline-blue-500 mb-6">
        <Search className="w-6 h-6 text-zinc-600 mr-2" />
        <input
          placeholder="Search Todo"
          type="text"
          autoComplete="off"
          onChange={onSearch}
          className="text-black focus:outline-none focus:ring-0 focus:border-none h-8 w-full"
        />
      </div>
      <ul className="px-4">
        <li className="-mx-4 mb-2 hover:cursor-pointer">
          <Button
            variant="text"
            isBlock
            className={cn(btnClassName, search.search.tag === '' ? 'bg-[#EDEBEB]' : 'bg-transparent')}
            onClick={() => search.setTag('')}
          >
            <List className="w-5 h-5 text-black" />
            <span className="text-[16px]">All Todo</span>
          </Button>
        </li>
        <li className="-mx-4 mb-2 hover:cursor-pointer">
          <Button
            variant="text"
            isBlock
            className={cn(btnClassName, search.search.tag === 'active' ? 'bg-[#EDEBEB]' : 'bg-transparent')}
            onClick={() => search.setTag('active')}
          >
            <LayoutList className="w-5 h-5 text-red-600" />
            <span className="text-[16px]">Not Finished Todo</span>
          </Button>
        </li>
        <li className="-mx-4 hover:cursor-pointer">
          <Button
            variant="text"
            isBlock
            className={cn(btnClassName, search.search.tag === 'inactive' ? 'bg-[#EDEBEB]' : 'bg-transparent')}
            onClick={() => search.setTag('inactive')}
          >
            <ListChecks className="w-5 h-5 text-blue-700" />
            <span className="text-[16px]">Finished Todo</span>
          </Button>
        </li>
      </ul>
      <Button variant="danger" className="flex gap-2 mt-4 items-center self-center" onClick={handleLogout}>
        <LogOut className="w-5 h-5" />
        Logout
      </Button>
    </div>
  );
};

export default Profile;
