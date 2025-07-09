export interface UserState {
  user: {
    id: string;
    email: string;
    name: string;
  } | null;
  isLoading: boolean;
  setUserLoading: (v: boolean) => void;
  setUser: (v: UserState['user']) => void;
}

export interface AuthState {
  isAuthenticate: boolean;
  revalidate: number;
  lastFetch: number;
  setIsAuthenticate: (v: boolean) => void;
  setLastFetch: (v: number) => void;
}
