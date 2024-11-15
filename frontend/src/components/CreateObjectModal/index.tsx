import React, { ChangeEvent, useCallback, useState } from 'react';
import { Button, Checkbox, Flex, Input, Modal, Title } from '@mantine/core';

import { useCreateMutation } from '../../slices/Object';
import { Object } from '../../models/Object';

export interface CreateObjectModalProps {
  opened: boolean;
  onClose(): void;
}

export const CreateObjectModal = ({
  opened,
  onClose,
}: CreateObjectModalProps) => {
  const [name, setName] = useState<string>('');
  const [isTemplate, setIsTemplate] = useState<boolean>(false);
  const [templateQuery, setTemplateQuery] = useState<string>('');

  const [createObject] = useCreateMutation();

  const nameUpdated = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setName(evt.target.value);
  }, []);

  const isTemplateToggled = useCallback(() => {
    if (isTemplate) {
      setIsTemplate(false);
    } else {
      setIsTemplate(true);
    }
  }, [isTemplate]);

  const templateQueryUpdated = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      setTemplateQuery(evt.target.value);
    },
    [],
  );

  const createClicked = useCallback(() => {
    if (name) {
      createObject({
        name,
        isTemplate,
      }).then(res => {
        onClose();
        window.location.href = `/#/o/${(res as { data: Object }).data.id}`;
      });
    }
  }, [createObject, isTemplate, name, onClose]);

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
        <Checkbox
          label="Template?"
          checked={isTemplate}
          onChange={isTemplateToggled}
        />
        <Input
          name={templateQuery}
          onChange={templateQueryUpdated}
          size="xs"
          placeholder="Search for template"
        />
        <Button onClick={createClicked}>Create</Button>
      </Flex>
    </Modal>
  );
};
