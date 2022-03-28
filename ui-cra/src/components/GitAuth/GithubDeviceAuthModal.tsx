import { CircularProgress } from '@material-ui/core';
import {
  Alert,
  Button,
  Flex,
  Icon,
  IconType,
  Modal,
  Text,
  useAuth,
} from '@weaveworks/weave-gitops';
import * as React from 'react';
import styled from 'styled-components';
import { GithubAuthContext } from '../../contexts/GitAuth/GithubAuthContext';
import {
  GitProvider,
  GetGithubDeviceCodeResponse,
} from '@weaveworks/weave-gitops/ui/lib/api/applications/applications.pb';

type Props = {
  className?: string;
  bodyClassName?: string;
  open: boolean;
  onSuccess: (token: string) => void;
  onClose: () => void;
  repoName: string;
};

const Pad = styled(Flex)`
  padding: 8px 0;
`;

const ModalContent = styled(({ codeRes, onSuccess, onError, className }) => {
  // Move this to a component so that we get the cancel logic when the modal closes.
  const { getGithubAuthStatus } = useAuth();
  const [loading, setLoading] = React.useState(true);

  React.useEffect(() => {
    if (!codeRes) {
      return;
    }
    setLoading(true);
    const { cancel, promise } = getGithubAuthStatus(codeRes);

    promise
      .then((authRes: any) => {
        onSuccess(authRes.accessToken);
      })
      .catch((err: any) => {
        onError(err);
      })
      .finally(() => setLoading(false));

    return cancel;
  }, [codeRes]);

  return (
    <div className={className}>
      <Pad wide center>
        {/* @ts-ignore */}
        <Text size="extraLarge">{codeRes.userCode}</Text>
      </Pad>
      <Pad wide center>
        <a target="_blank" href={codeRes.validationURI}>
          <Button
            type="button"
            startIcon={<Icon size="base" type={IconType.ExternalTab} />}
          >
            Authorize Github Access
          </Button>
        </a>
      </Pad>
      <Pad wide center>
        {loading && <div>Waiting for authorization to be completed...</div>}
      </Pad>
    </div>
  );
})`
  ${Icon} {
    margin-left: 8px;
  }
`;

function GithubDeviceAuthModal({
  className,
  bodyClassName,
  open,
  onClose,
  repoName,
  onSuccess,
}: Props) {
  const [codeRes, setCodeRes] =
    // @ts-ignore
    React.useState<GetGithubDeviceCodeResponse>(null);
  const { getGithubDeviceCode, storeProviderToken } = useAuth();
  const [codeLoading, setCodeLoading] = React.useState(true);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    if (!open) {
      return;
    }

    setCodeLoading(true);

    getGithubDeviceCode()
      .then((res: any) => {
        setCodeRes(res);
      })
      .finally(() => setCodeLoading(false));
  }, [open]);
  return (
    <Modal
      className={className}
      bodyClassName={bodyClassName}
      title="Authenticate with Github"
      open={open}
      onClose={onClose}
      description={`Weave GitOps needs to authenticate with the Git Provider for the ${repoName} repo`}
    >
      <p>
        Paste this code into the Github Device Activation field to grant Weave
        GitOps temporary access:
      </p>
      {error && (
        // @ts-ignore
        <Alert severity="error" title="Error" message={error.message} center />
      )}
      <Flex wide center height="150px">
        {codeLoading || !codeRes ? (
          <CircularProgress />
        ) : (
          <ModalContent
            onSuccess={(token: any) => {
              storeProviderToken(GitProvider.GitHub, token);
              onSuccess(token);
              onClose();
            }}
            onError={(err: any) => setError(err)}
            codeRes={codeRes}
          />
        )}
      </Flex>
    </Modal>
  );
}

const StyledModal = styled(GithubDeviceAuthModal)``;

export default StyledModal;

export function GlobalGithubAuthDialog() {
  const { dialogState, setDialogState, setSuccess } =
    React.useContext(GithubAuthContext);

  return (
    <StyledModal
      repoName={dialogState.repoName}
      open={dialogState.open}
      onClose={() => setDialogState(false, dialogState.repoName)}
      onSuccess={setSuccess}
    />
  );
}
