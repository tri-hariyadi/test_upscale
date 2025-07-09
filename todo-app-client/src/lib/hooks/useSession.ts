import { useEffect } from 'react';

import api from 'lib/api/api.ts';
import { useUserStore } from 'store/AuthStore';

const useSession = () => {
  const userStore = useUserStore();
  const token = window.localStorage.getItem('access_token');

  useEffect(() => {
    const verify = async () => {
      userStore.setUserLoading(true);
      const user = await api.verifyToken({ headers: { Cookie: `access_token=${token}` } });
      if (user.data) {
        userStore.setUser(user.data);
      }
      userStore.setUserLoading(false);
    };

    verify();
  }, []);

  return userStore;
};

export default useSession;
