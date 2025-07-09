import DummyProfile from 'assets/dummy_profile.png';

import Dropdown, { DropdownMenuContent, DropdownMenuTrigger } from '../Dropdown';
import Profile from '../Profile';

const Header = () => {
  return (
    <div className="h-16 w-full bg-white py-3 flex justify-center shadow-md fixed top-0 z-50">
      <div className="xl:min-w-5xl min-w-full xl:max-w-5xl px-4 sm:px-8 flex justify-between items-center">
        <h2 className="text-blue-700 font-inter-extra-bold text-2xl">Todo App</h2>
        <div className="flex items-center gap-3">
          <div className="block md:hidden">
            <Dropdown showOverlay={false}>
              <DropdownMenuTrigger>
                <button>
                  <img src={DummyProfile} alt="profile" className="object-fill w-10 h-10" />
                </button>
              </DropdownMenuTrigger>
              <DropdownMenuContent className="px-4 py-6">
                <Profile />
              </DropdownMenuContent>
            </Dropdown>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
