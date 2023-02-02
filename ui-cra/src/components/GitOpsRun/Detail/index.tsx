import { RouterTab, SubRouterTabs } from '@weaveworks/weave-gitops';
import { useRouteMatch } from 'react-router-dom';
import styled from 'styled-components';
import { ContentWrapper } from '../../Layout/ContentWrapper';
import { PageTemplate } from '../../Layout/PageTemplate';
import GitOpsRunLogs from './GitOpsRunLogs';
type Props = {
  name?: string;
  namespace?: string;
};

const PageTitle = styled.h4`
  font-size: ${({ theme }) => theme.fontSizes.large};
  font-weight: 600;
  margin: 0;
  margin-bottom: ${({ theme }) => theme.spacing.small};
`;

const GitOpsRunDetail = ({ name, namespace }: Props) => {
  const { path } = useRouteMatch();
  return (
    <PageTemplate
      documentTitle="GitOps Run Detail"
      path={[{ label: 'GitOps Run Detail' }]}
    >
      <ContentWrapper>
        <PageTitle>{name}</PageTitle>
        <SubRouterTabs rootPath={`${path}/logs`}>
          <RouterTab name="Logs" path={`${path}/logs`}>
            <GitOpsRunLogs name={name || ''} namespace={namespace || ''} />
          </RouterTab>
        </SubRouterTabs>
      </ContentWrapper>
    </PageTemplate>
  );
};

export default GitOpsRunDetail;
