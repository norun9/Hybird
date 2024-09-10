FROM amazon/aws-cli

ENV AWS_ACCESS_KEY_ID=fake_access_key\
    AWS_SECRET_ACCESS_KEY=fake_secret_access_key\
    DYNAMODB_REGION=ap-northeast-1

RUN mkdir -p /root/.aws && \
    echo "[default]" > /root/.aws/credentials && \
    echo "aws_access_key_id = ${AWS_ACCESS_KEY_ID}" >> /root/.aws/credentials && \
    echo "aws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}" >> /root/.aws/credentials && \
    echo "[default]" > /root/.aws/config && \
    echo "region = ap-northeast-1" >> /root/.aws/config

WORKDIR /usr/app