import React, { useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { Container, Flex, LoadingOverlay, Table, Title } from '@mantine/core';

import { useShowQuery } from '../../slices/Object';
import { Object } from '../../models/Object';

export const ObjectPage = () => {
  const { id } = useParams();

  const { data: objectData, isLoading: objectLoading } = useShowQuery(id || '');
  const loading = useMemo(() => objectLoading, [objectLoading]);

  const object = useMemo(() => {
    if (objectData) {
      return new Object(objectData);
    }

    return null;
  }, [objectData]);

  return (
    <Container size="lg" style={{ flex: 1 }}>
      <Flex direction="column" align="center">
        <Title my="sm">{object?.name}</Title>
        <Table withTableBorder bg={object?.isTemplate ? '#ffffdd' : 'blue'}>
          {object?.isTemplate && (
            <Table.Thead>
              <Table.Tr>
                <Table.Td colSpan={2} align="center">
                  <Title order={6} tt="uppercase">
                    Template
                  </Title>
                </Table.Td>
              </Table.Tr>
            </Table.Thead>
          )}
          <Table.Tbody>
            <Table.Tr>
              <Table.Th>Name</Table.Th>
              <Table.Td>{object?.name}</Table.Td>
            </Table.Tr>
            <Table.Tr>
              <Table.Th>ID</Table.Th>
              <Table.Td>{object?.id}</Table.Td>
            </Table.Tr>
          </Table.Tbody>
        </Table>
      </Flex>
      <LoadingOverlay visible={loading} />
    </Container>
  );
};
