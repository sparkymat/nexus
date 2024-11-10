/* eslint-disable react/no-unknown-property */
import React, { useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { Card, Container, Text } from '@mantine/core';

export interface SearchResultsProps {}

// eslint-disable-next-line no-empty-pattern
export const SearchResults = ({}: SearchResultsProps) => {
  const { query: queryHash } = useParams();

  const query = useMemo(() => {
    if (queryHash) {
      return atob(queryHash);
    }

    return '';
  }, [queryHash]);

  return (
    <Container
      py="md"
      style={{
        flex: 1,
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <Card p="xs">
        <Text>{`Searching for '${query}'`}</Text>
      </Card>
    </Container>
  );
};
