export interface SearchState {
  search: {
    tag: string;
    keyword: string;
  };
  setTag: (tag: string) => void;
  setKeyword: (keyword: string) => void;
}
