/* eslint-disable react/no-unknown-property */
import React, { ChangeEvent, useCallback, useMemo, useState } from 'react';
import { Button, Container, Flex, Input } from '@mantine/core';

export interface HomeProps {}

// eslint-disable-next-line no-empty-pattern
export const Home = ({}: HomeProps) => {
  const [query, setQuery] = useState<string>('');

  const queryUpdated = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setQuery(evt.target.value);
  }, []);

  const canSubmitQuery = useMemo(() => !!query, [query]);

  const querySubmitted = useCallback(() => {
    if (canSubmitQuery) {
      const queryHash = btoa(query);

      window.location.href = `/#/q/${queryHash}`;
    }
  }, [canSubmitQuery, query]);

  const queryKeyUp = useCallback(
    (evt: React.KeyboardEvent) => {
      if (evt.keyCode === 13 && canSubmitQuery) {
        const queryHash = btoa(query);

        window.location.href = `/#/q/${queryHash}`;
      }
    },
    [canSubmitQuery, query],
  );

  return (
    <Container
      py="md"
      style={{
        flex: 1,
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
      }}
    >
      <Flex gap="sm" wrap="wrap">
        <Input
          autoFocus
          placeholder="Enter your search query here"
          value={query}
          onChange={queryUpdated}
          onKeyUp={queryKeyUp}
          size="lg"
          style={{ flex: 1 }}
        />
        <Button
          disabled={!canSubmitQuery}
          h="100%"
          size="compact-md"
          onClick={querySubmitted}
        >
          Search
        </Button>
      </Flex>
    </Container>
  );
};
