/* eslint-disable react/no-unknown-property */
import React, { useCallback, useEffect, useState } from 'react';
import { Routes, Route, useLocation } from 'react-router-dom';
import {
  ActionIcon,
  Anchor,
  AppShell,
  Burger,
  Flex,
  LoadingOverlay,
  Menu,
  Space,
  Title,
  useComputedColorScheme,
  useMantineColorScheme,
  useMantineTheme,
} from '@mantine/core';
import { useDispatch } from 'react-redux';
import {
  IconBrightness,
  IconFilePlus,
  IconUserCircle,
} from '@tabler/icons-react';
import { useDisclosure } from '@mantine/hooks';

import { AppDispatch } from '../../store';
import { updatePath } from '../../slices/App';
import { Home } from '../Home';
import { SearchResults } from '../SearchResults';
import { CreateObjectModal } from '../CreateObjectModal';
import { ObjectPage } from '../ObjectPage';

// Start of app import code generated by oxgen. DO NOT EDIT.
// End of app import code generated by oxgen. DO NOT EDIT.

export const App = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [opened, { toggle }] = useDisclosure();
  const location = useLocation();
  const theme = useMantineTheme();

  const { toggleColorScheme } = useMantineColorScheme();
  const computedColorScheme = useComputedColorScheme('light', {
    getInitialValueInEffect: true,
  });
  const [createModalOpen, setCreateModalOpen] = useState<boolean>(false);

  const createModalOpened = useCallback(() => {
    setCreateModalOpen(true);
  }, []);

  const createModalClosed = useCallback(() => {
    setCreateModalOpen(false);
  }, []);

  useEffect(() => {
    dispatch(updatePath(location.pathname));
  }, [dispatch, location]);

  return (
    <div>
      <AppShell
        header={{ height: 40 }}
        styles={thisTheme => ({
          main: {
            backgroundColor:
              computedColorScheme === 'dark'
                ? thisTheme.colors.dark[8]
                : thisTheme.colors.gray[0],
          },
        })}
      >
        <AppShell.Header
          h={40}
          px={{ base: 'xs', md: 'md' }}
          py={{ base: '0px', md: 'md' }}
        >
          <div
            style={{ display: 'flex', alignItems: 'center', height: '100%' }}
          >
            <Flex
              align="center"
              direction="row"
              wrap="wrap"
              style={{ width: '100%' }}
            >
              <Burger
                opened={opened}
                onClick={toggle}
                hiddenFrom="sm"
                size="xs"
                color={theme.colors.gray[6]}
                mr="xl"
              />
              <Anchor
                c={computedColorScheme === 'dark' ? 'white' : 'dark'}
                href="/#/"
              >
                <Title order={3}>nexus</Title>
              </Anchor>

              <Space style={{ flex: 1 }} w="sm" />

              <Menu shadow="md" width={200}>
                <Menu.Target>
                  <ActionIcon
                    c={computedColorScheme === 'dark' ? 'white' : 'dark'}
                  >
                    <IconUserCircle strokeWidth={1} size={24} />
                  </ActionIcon>
                </Menu.Target>

                <Menu.Dropdown>
                  <Menu.Item
                    leftSection={<IconFilePlus strokeWidth={1} size={14} />}
                    onClick={() => createModalOpened()}
                  >
                    New object
                  </Menu.Item>
                  <Menu.Item
                    leftSection={<IconBrightness strokeWidth={1} size={14} />}
                    onClick={() => toggleColorScheme()}
                  >
                    Toggle dark mode
                  </Menu.Item>
                </Menu.Dropdown>
              </Menu>
            </Flex>
          </div>
        </AppShell.Header>
        <AppShell.Main style={{ display: 'flex' }}>
          <Routes>
            <Route index element={<Home />} />
            {/* Start of app route code generated by oxgen. DO NOT EDIT. */}
            <Route path="/q/:query" element={<SearchResults />} />
            <Route path="/o/:id" element={<ObjectPage />} />
            {/* End of app route code generated by oxgen. DO NOT EDIT. */}
          </Routes>
        </AppShell.Main>
      </AppShell>
      <CreateObjectModal opened={createModalOpen} onClose={createModalClosed} />
      <LoadingOverlay visible={false} />
    </div>
  );
};
