import { Box, Flex, Heading } from '@radix-ui/themes';

import { HamburgerMenuButton } from 'components';
import { Header, PageLayout } from 'components/layout';
import { useOllamaModels } from 'hooks';

export const Settings: React.FC = () => {
  const { data: models } = useOllamaModels();

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular">
          Models
        </Heading>
        <Box width="40px"></Box>
      </Header>
      <Flex direction="column" align="center" justify="center" flexGrow="1">
        <Box>
          {models?.map((model) => (
            <Box key={model.id}>
              {model.name}:{model.tag}
            </Box>
          ))}
        </Box>
      </Flex>
    </PageLayout>
  );
};
