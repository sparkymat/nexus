import React, { ChangeEvent, useCallback, useState } from 'react';
import { Button, Checkbox, Flex, Input, Modal, Title } from '@mantine/core';

export interface CreateObjectModalProps {
  opened: boolean;
  onClose(): void;
}

export const CreateObjectModal = ({
  opened,
  onClose,
}: CreateObjectModalProps) => {
  const [name, setName] = useState<string>('');
  const [templateQuery, setTemplateQuery] = useState<string>('');

  const nameUpdated = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setName(evt.target.value);
  }, []);

  const templateQueryUpdated = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      setTemplateQuery(evt.target.value);
    },
    [],
  );

  return (
    <Modal opened={opened} onClose={onClose} withCloseButton={false}>
      <Flex direction="column" align="stretch" gap="md">
        <Title order={3}>new object</Title>
        <Input
          name={name}
          onChange={nameUpdated}
          size="md"
          placeholder="Name of the object"
        />
        <Checkbox label="Template?" />
        <Input
          name={templateQuery}
          onChange={templateQueryUpdated}
          size="xs"
          placeholder="Search for template"
        />
        <Button>Create</Button>
      </Flex>
    </Modal>
  );
};
