import { create } from 'zustand/react';

import type { SearchState } from 'store/SearchStore/types';

export const useSearchStore = create<SearchState>((set, getState) => ({
  search: { tag: '', keyword: '' },
  setTag: (tag) => set({ search: { ...getState().search, tag } }),
  setKeyword: (keyword) => set({ search: { ...getState().search, keyword } })
}));
