import { useEffect, useState } from 'react';

import { Fallback, Show } from 'components';
import api from 'lib/api/api.ts';
import type { IApi } from 'lib/api/types';
import { wrapPromise } from 'lib/utils.ts';

const getQuote = async () => {
  const resp = await api.getQuote();
  if (resp.status) {
    return resp.data;
  }
};

const Quote = () => {
  const [resource, setResource] = useState<{
    read(): IApi['Quote'] | undefined;
  } | null>();

  useEffect(() => {
    const promise = wrapPromise(getQuote());
    setResource(promise);
  }, []);

  if (!resource) return <Fallback />;

  const quote = resource?.read();
  // const quote = {};
  // Test CI/CD
  return (
    <blockquote className="border-l-4 border-[#78C0A8] py-[1.2em] pr-8 pl-14 font-inter-italic text-gray-700 my-4 relative bg-[#EDEDED] before:content-['“'] before:text-[#78C0A8] before:text-[4em] before:absolute before:left-3 before:-top-3">
      <Show>
        <Show.When isTrue={quote}>
          {quote?.quote}
          <span className="block text-[#333333] font-inter-bold-italic mt-4">{quote?.author}</span>
        </Show.When>
        <Show.Else>
          <span className="text-neutral-500 font-inter-italic">No quotes today</span>
        </Show.Else>
      </Show>
    </blockquote>
  );
};

export default Quote;
