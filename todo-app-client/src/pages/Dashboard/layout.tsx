import type { PropsWithChildren } from 'react';

import { cn, getChildByType } from 'lib/utils.ts';

export const LayoutLeft = ({ children, className }: PropsWithChildren<{ className?: string }>) => (
  <div className={cn('sticky top-28 bg-white', className)}>{children}</div>
);
export const LayoutRight = ({ children }: PropsWithChildren) => children;

const Layout = ({ children }: PropsWithChildren) => {
  const item = getChildByType(children, LayoutRight);
  const menu = getChildByType(children, LayoutLeft);

  return (
    <div className="flex justify-center pb-12">
      <div className="grid grid-cols-16 gap-0 md:gap-8 mt-28 px-4 sm:px-8 xl:min-w-5xl min-w-full xl:max-w-5xl">
        <div className="col-span-4 md:col-span-7 lg:col-span-6 hidden md:block">{menu}</div>

        <div className="col-span-16 md:col-span-9 lg:col-span-10">{item}</div>
      </div>
    </div>
  );
};

export default Layout;
