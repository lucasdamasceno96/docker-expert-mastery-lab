# Docker Cheatsheet Completo

Este guia reúne os comandos mais importantes do Docker, organizados por tema, com exemplos e o melhor momento para usá-los. A ideia é servir como referência rápida para estudo, laboratório e uso no dia a dia.

---

## 1) Visão geral da CLI

### `docker --help`

Mostra os comandos disponíveis e a ajuda geral.

```bash
docker --help
```

**Use quando:** você quer descobrir subcomandos, opções e formato básico.

### `docker <comando> --help`

Exibe a ajuda de um comando específico.

```bash
docker run --help
docker build --help
```

**Use quando:** você quer lembrar flags, modos e opções avançadas de um comando.

### `docker version`

Mostra a versão do cliente e do daemon.

```bash
docker version
```

**Use quando:** precisa validar compatibilidade entre CLI e engine.

### `docker info`

Mostra um resumo do ambiente Docker: storage driver, cgroups, redes, volumes, plugins, etc.

```bash
docker info
```

**Use quando:** quer diagnosticar o estado da instalação ou entender como o host está configurado.

---

## 2) Imagens

### `docker images`

Lista as imagens locais.

```bash
docker images
```

**Use quando:** quer ver o que já está baixado ou construído.

### `docker pull`

Baixa uma imagem do registry.

```bash
docker pull nginx:latest
```

**Use quando:** precisa obter uma imagem antes de rodar ou testar.

### `docker push`

Envia uma imagem para um registry.

```bash
docker push meuuser/minha-app:1.0.0
```

**Use quando:** quer publicar uma imagem para CI/CD, Kubernetes ou compartilhamento.

### `docker build`

Constrói uma imagem a partir de um `Dockerfile`.

```bash
docker build -t minha-app:1.0.0 .
```

**Use quando:** está criando uma imagem customizada.

### `docker build --no-cache`

Força a recompilação de todas as camadas.

```bash
docker build --no-cache -t minha-app:debug .
```

**Use quando:** suspeita de cache inválido ou quer garantir build limpo.

### `docker build --target`

Constrói até um estágio específico de multi-stage build.

```bash
docker build --target builder -t minha-app:builder .
```

**Use quando:** quer testar uma etapa intermediária ou depurar o build.

### `docker build -f`

Define explicitamente o arquivo Dockerfile.

```bash
docker build -f Dockerfile.prod -t minha-app:prod .
```

**Use quando:** o Dockerfile não está no nome/padrão padrão.

### `docker tag`

Cria outra tag para a mesma imagem.

```bash
docker tag minha-app:1.0.0 meuuser/minha-app:latest
```

**Use quando:** quer versionar, publicar ou renomear uma imagem.

### `docker rmi`

Remove uma ou mais imagens.

```bash
docker rmi minha-app:1.0.0
```

**Use quando:** quer limpar imagens antigas ou liberar espaço.

### `docker history`

Mostra o histórico de camadas da imagem.

```bash
docker history minha-app:1.0.0
```

**Use quando:** quer entender o tamanho de cada camada e otimizar o build.

### `docker save`

Exporta uma imagem para um arquivo `.tar`.

```bash
docker save -o nginx.tar nginx:latest
```

**Use quando:** precisa transportar uma imagem sem registry.

### `docker load`

Importa uma imagem de um arquivo `.tar` gerado por `docker save`.

```bash
docker load -i nginx.tar
```

**Use quando:** quer restaurar uma imagem offline.

### `docker import`

Cria uma imagem a partir de um arquivo `.tar` de filesystem, sem histórico de camadas.

```bash
docker import rootfs.tar minha-imagem:base
```

**Use quando:** quer transformar um root filesystem em imagem rapidamente.

---

## 3) Containers

### `docker run`

Cria e inicia um container em um único passo.

```bash
docker run --name web -p 8080:80 nginx:latest
```

**Use quando:** quer subir um container rapidamente.

### `docker run -d`

Executa em segundo plano.

```bash
docker run -d --name web nginx:latest
```

**Use quando:** o container é um serviço de longa duração.

### `docker run --rm`

Remove o container automaticamente ao terminar.

```bash
docker run --rm alpine echo "ok"
```

**Use quando:** faz testes descartáveis.

### `docker run -it`

Modo interativo com TTY.

```bash
docker run -it ubuntu bash
```

**Use quando:** precisa de shell dentro do container.

### `docker run --name`

Define um nome fixo para o container.

```bash
docker run -d --name api minha-api:1.0.0
```

**Use quando:** quer facilitar referência, logs e troubleshooting.

### `docker run -p`

Publica portas do container para o host.

```bash
docker run -d -p 8080:80 nginx
```

**Use quando:** precisa acessar o serviço a partir do host ou da rede externa.

### `docker run -P`

Publica automaticamente portas expostas.

```bash
docker run -d -P nginx
```

**Use quando:** quer evitar escolher portas manualmente em testes rápidos.

### `docker run -e`

Define variáveis de ambiente.

```bash
docker run -e ENV=prod -e DEBUG=false minha-app
```

**Use quando:** precisa parametrizar o app.

### `docker run --env-file`

Carrega variáveis de ambiente de um arquivo.

```bash
docker run --env-file .env minha-app
```

**Use quando:** quer separar configuração do comando.

### `docker run -v`

Monta volumes ou bind mounts.

```bash
docker run -v meus-dados:/data alpine
```

**Use quando:** precisa persistência ou compartilhar arquivos com o host.

### `docker run --mount`

Forma mais explícita e moderna de montar volumes, binds e tmpfs.

```bash
docker run --mount type=volume,src=meus-dados,dst=/data alpine
```

**Use quando:** quer clareza e controle fino sobre o tipo de mount.

### `docker run --network`

Conecta o container a uma rede específica.

```bash
docker run --network minha-rede minha-app
```

**Use quando:** precisa de comunicação entre containers.

### `docker run --hostname`

Define o hostname interno do container.

```bash
docker run --hostname app01 minha-app
```

**Use quando:** o hostname importa para logs, testes ou integração.

### `docker run --restart`

Configura política de reinício.

```bash
docker run -d --restart unless-stopped minha-app
```

**Use quando:** quer comportamento resiliente em produção.

### `docker run --memory`

Limita memória.

```bash
docker run --memory 512m minha-app
```

**Use quando:** quer testar ou impor limites de recursos.

### `docker run --cpus`

Limita CPU.

```bash
docker run --cpus 1.5 minha-app
```

**Use quando:** quer controlar consumo de CPU.

### `docker run --cpuset-cpus`

Fixa CPUs específicas.

```bash
docker run --cpuset-cpus="0,1" minha-app
```

**Use quando:** precisa isolar processamento em núcleos específicos.

### `docker run --read-only`

Deixa o filesystem do container somente leitura.

```bash
docker run --read-only minha-app
```

**Use quando:** quer endurecimento de segurança.

### `docker run --user`

Executa com um usuário específico.

```bash
docker run --user 1000:1000 minha-app
```

**Use quando:** quer evitar rodar como root.

### `docker run --cap-drop`

Remove capabilities do Linux.

```bash
docker run --cap-drop ALL --cap-add NET_BIND_SERVICE minha-app
```

**Use quando:** quer reduzir privilégios do container.

### `docker run --security-opt`

Aplica opções de segurança, como seccomp, AppArmor ou no-new-privileges.

```bash
docker run --security-opt no-new-privileges:true minha-app
```

**Use quando:** precisa endurecer a execução.

### `docker run --privileged`

Executa com privilégios elevados.

```bash
docker run --privileged ubuntu
```

**Use quando:** apenas para debug ou cenários muito específicos. Evite em produção.

### `docker ps`

Lista containers em execução.

```bash
docker ps
```

**Use quando:** quer ver o que está rodando.

### `docker ps -a`

Lista todos os containers, inclusive parados.

```bash
docker ps -a
```

**Use quando:** quer investigar falhas ou containers encerrados.

### `docker stop`

Envia SIGTERM e depois SIGKILL se necessário.

```bash
docker stop web
```

**Use quando:** quer encerrar o container de forma limpa.

### `docker start`

Inicia um container parado.

```bash
docker start web
```

**Use quando:** quer reativar um container já criado.

### `docker restart`

Reinicia um container.

```bash
docker restart web
```

**Use quando:** quer aplicar nova execução sem recriar.

### `docker kill`

Encerra imediatamente, por padrão com SIGKILL.

```bash
docker kill web
```

**Use quando:** o container não responde ao stop.

### `docker pause`

Suspende processos do container.

```bash
docker pause web
```

**Use quando:** quer congelar execução temporariamente.

### `docker unpause`

Retoma um container pausado.

```bash
docker unpause web
```

**Use quando:** quer continuar a execução.

### `docker rm`

Remove containers parados.

```bash
docker rm web
```

**Use quando:** quer limpar containers antigos.

### `docker rm -f`

Força remoção de container em execução.

```bash
docker rm -f web
```

**Use quando:** quer remover imediatamente, inclusive rodando.

### `docker rename`

Renomeia um container.

```bash
docker rename web web-old
```

**Use quando:** quer padronizar nomes.

### `docker exec`

Executa um comando dentro de um container em execução.

```bash
docker exec -it web sh
```

**Use quando:** precisa inspecionar ou debugar um container vivo.

### `docker attach`

Conecta ao processo principal de um container.

```bash
docker attach web
```

**Use quando:** quer ver a saída do processo principal diretamente.

### `docker cp`

Copia arquivos entre host e container.

```bash
docker cp arquivo.txt web:/tmp/arquivo.txt
```

**Use quando:** precisa transferir arquivos para debug ou análise.

### `docker diff`

Mostra diferenças no filesystem do container em relação à imagem base.

```bash
docker diff web
```

**Use quando:** quer entender o que foi alterado dentro do container.

---

## 4) Logs e inspeção

### `docker logs`

Mostra logs do container.

```bash
docker logs web
```

**Use quando:** quer diagnosticar comportamento da aplicação.

### `docker logs -f`

Segue os logs em tempo real.

```bash
docker logs -f web
```

**Use quando:** está observando inicialização ou reproduzindo um problema.

### `docker logs --tail`

Mostra as últimas linhas.

```bash
docker logs --tail 100 web
```

**Use quando:** quer contexto recente sem imprimir tudo.

### `docker inspect`

Mostra JSON detalhado de container, imagem, volume, rede ou outros objetos.

```bash
docker inspect web
```

**Use quando:** precisa de informações técnicas profundas.

### `docker stats`

Exibe consumo de CPU, memória, I/O e rede em tempo real.

```bash
docker stats
```

**Use quando:** quer avaliar desempenho e consumo.

### `docker top`

Mostra processos rodando dentro do container.

```bash
docker top web
```

**Use quando:** quer ver a árvore de processos ou detectar processos extras.

### `docker events`

Mostra eventos do daemon em tempo real.

```bash
docker events
```

**Use quando:** quer observar criação, remoção, rede, volume e lifecycle.

### `docker wait`

Bloqueia até o container terminar e retorna o exit code.

```bash
docker wait web
```

**Use quando:** precisa automatizar fluxos e capturar status final.

---

## 5) Redes

### `docker network ls`

Lista redes existentes.

```bash
docker network ls
```

**Use quando:** quer ver redes criadas pelo Docker.

### `docker network create`

Cria uma rede.

```bash
docker network create minha-rede
```

**Use quando:** quer isolar containers em uma rede própria.

### `docker network inspect`

Mostra detalhes da rede.

```bash
docker network inspect minha-rede
```

**Use quando:** quer ver IPs, containers conectados, subnets e gateway.

### `docker network connect`

Conecta um container a uma rede.

```bash
docker network connect minha-rede web
```

**Use quando:** um container precisa participar de outra rede.

### `docker network disconnect`

Remove um container de uma rede.

```bash
docker network disconnect minha-rede web
```

**Use quando:** quer reorganizar conectividade.

### `docker network rm`

Remove uma rede.

```bash
docker network rm minha-rede
```

**Use quando:** a rede não é mais necessária.

**Casos de uso comuns de rede:**

- comunicação entre serviços por nome DNS interno;
- isolamento entre ambientes;
- depuração de conectividade;
- criação de redes com sub-redes específicas e IP fixo via IPAM.

---

## 6) Volumes e armazenamento

### `docker volume ls`

Lista volumes.

```bash
docker volume ls
```

**Use quando:** quer ver persistência disponível no host.

### `docker volume create`

Cria um volume.

```bash
docker volume create meus-dados
```

**Use quando:** precisa de persistência desacoplada do container.

### `docker volume inspect`

Mostra detalhes do volume.

```bash
docker volume inspect meus-dados
```

**Use quando:** quer descobrir o caminho físico e metadados.

### `docker volume rm`

Remove um volume.

```bash
docker volume rm meus-dados
```

**Use quando:** o volume não é mais necessário.

### `docker volume prune`

Remove volumes não utilizados.

```bash
docker volume prune
```

**Use quando:** quer limpar espaço com segurança relativa.

### `docker run -v /host:/container`

Bind mount entre host e container.

```bash
docker run -v $(pwd)/data:/app/data alpine
```

**Use quando:** quer compartilhar arquivos do host ou usar código local.

### `docker run --mount type=bind`

Forma explícita para bind mount.

```bash
docker run --mount type=bind,src=$(pwd)/data,dst=/app/data alpine
```

**Use quando:** quer sintaxe mais legível e precisa.

### `docker run --mount type=tmpfs`

Cria storage temporário em memória.

```bash
docker run --mount type=tmpfs,dst=/tmp alpine
```

**Use quando:** precisa de dados temporários e rápidos.

**Quando usar volume x bind mount:**

- **Volume**: melhor para persistência gerenciada pelo Docker, backup e portabilidade.
- **Bind mount**: melhor para desenvolvimento local e acesso direto a arquivos do host.

---

## 7) Build avançado e BuildKit

### `DOCKER_BUILDKIT=1 docker build`

Ativa BuildKit no build.

```bash
DOCKER_BUILDKIT=1 docker build -t minha-app .
```

**Use quando:** quer builds mais rápidos, cache melhor e recursos avançados.

### `docker buildx ls`

Lista builders e ambientes buildx.

```bash
docker buildx ls
```

**Use quando:** quer verificar suporte a builds multiplataforma e BuildKit avançado.

### `docker buildx create`

Cria um builder.

```bash
docker buildx create --name meu-builder --use
```

**Use quando:** precisa de um ambiente isolado para build avançado.

### `docker buildx inspect`

Mostra detalhes do builder.

```bash
docker buildx inspect --bootstrap
```

**Use quando:** quer confirmar que o builder está pronto.

### `docker buildx build`

Build avançado com suporte a cache, export, múltiplas plataformas e mais.

```bash
docker buildx build -t minha-app:1.0.0 --load .
```

**Use quando:** quer recursos modernos do BuildKit.

### `--platform`

Define plataforma alvo.

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t minha-app:multi --push .
```

**Use quando:** precisa de imagem multi-arquitetura.

### `--secret`

Passa secrets de forma segura no build.

```bash
docker buildx build --secret id=npmrc,src=$HOME/.npmrc .
```

**Use quando:** precisa de credenciais sem gravá-las na imagem.

### `--ssh`

Encaminha agente SSH para o build.

```bash
docker buildx build --ssh default .
```

**Use quando:** precisa acessar repositórios privados durante o build.

### `--cache-from` / `--cache-to`

Importa e exporta cache de build.

```bash
docker buildx build --cache-from=type=registry,ref=meuuser/app:cache \
  --cache-to=type=registry,ref=meuuser/app:cache,mode=max \
  -t meuuser/app:1.0.0 .
```

**Use quando:** quer acelerar CI/CD e reaproveitar camadas.

### `--output`

Define saída do build.

```bash
docker buildx build --output type=local,dest=./out .
```

**Use quando:** quer gerar artefatos sem necessariamente criar uma imagem local.

### `--push`

Envia a imagem para o registry ao final do build.

```bash
docker buildx build --push -t meuuser/app:1.0.0 .
```

**Use quando:** a imagem será usada por pipeline ou cluster.

### `--load`

Carrega a imagem no Docker local.

```bash
docker buildx build --load -t minha-app:dev .
```

**Use quando:** quer testar imediatamente a imagem no host.

---

## 8) Docker Compose

> Dependendo da versão, o comando pode ser `docker compose` (moderno) em vez de `docker-compose` (legado).

### `docker compose up`

Cria e inicia os serviços.

```bash
docker compose up
```

**Use quando:** quer subir toda a stack definida no `compose.yaml`.

### `docker compose up -d`

Executa em segundo plano.

```bash
docker compose up -d
```

**Use quando:** quer trabalhar com a stack em background.

### `docker compose down`

Para e remove containers, redes e, opcionalmente, volumes.

```bash
docker compose down
```

**Use quando:** quer desmontar o ambiente.

### `docker compose down -v`

Remove também volumes.

```bash
docker compose down -v
```

**Use quando:** quer resetar totalmente o ambiente.

### `docker compose ps`

Lista serviços e estado atual.

```bash
docker compose ps
```

**Use quando:** quer ver o status da stack.

### `docker compose logs`

Mostra logs dos serviços.

```bash
docker compose logs -f
```

**Use quando:** quer depurar múltiplos serviços ao mesmo tempo.

### `docker compose exec`

Executa comando dentro de um serviço.

```bash
docker compose exec app sh
```

**Use quando:** precisa debugar um serviço da stack.

### `docker compose build`

Constrói as imagens dos serviços.

```bash
docker compose build
```

**Use quando:** quer rebuild da stack.

### `docker compose pull`

Baixa imagens declaradas.

```bash
docker compose pull
```

**Use quando:** quer atualizar dependências de imagem.

### `docker compose restart`

Reinicia serviços.

```bash
docker compose restart app
```

**Use quando:** quer aplicar uma reinicialização sem recriar tudo.

### `docker compose config`

Mostra a configuração final resolvida.

```bash
docker compose config
```

**Use quando:** quer validar variáveis, merges e interpolação.

---

## 9) Sistema, limpeza e manutenção

### `docker system df`

Mostra uso de espaço por imagens, containers, volumes e cache.

```bash
docker system df
```

**Use quando:** quer entender onde o espaço está sendo consumido.

### `docker system prune`

Remove recursos não utilizados.

```bash
docker system prune
```

**Use quando:** quer limpar lixo acumulado com cuidado.

### `docker system prune -a`

Remove imagens não usadas por containers.

```bash
docker system prune -a
```

**Use quando:** precisa de uma limpeza mais agressiva.

### `docker container prune`

Remove containers parados.

```bash
docker container prune
```

**Use quando:** quer limpar containers antigos sem mexer nas imagens.

### `docker image prune`

Remove imagens sem uso.

```bash
docker image prune
```

**Use quando:** quer liberar espaço em imagens dangling.

### `docker image prune -a`

Remove imagens não utilizadas por nenhum container.

```bash
docker image prune -a
```

**Use quando:** quer uma limpeza forte de imagens.

### `docker builder prune`

Remove cache de build.

```bash
docker builder prune
```

**Use quando:** o cache cresceu demais ou você quer forçar rebuild.

### `docker builder prune -a`

Remove praticamente todo cache não usado.

```bash
docker builder prune -a
```

**Use quando:** precisa de limpeza agressiva em pipelines ou hosts de build.

### `docker volume prune`

Remove volumes não utilizados.

```bash
docker volume prune
```

**Use quando:** quer limpar persistências órfãs.

### `docker network prune`

Remove redes não utilizadas.

```bash
docker network prune
```

**Use quando:** quer organizar redes antigas.

---

## 10) Contexts e acesso remoto

### `docker context ls`

Lista contexts disponíveis.

```bash
docker context ls
```

**Use quando:** trabalha com múltiplos ambientes Docker.

### `docker context use`

Troca o context ativo.

```bash
docker context use prod
```

**Use quando:** alterna entre local, remoto ou Docker Desktop.

### `docker context create`

Cria um novo context.

```bash
docker context create meu-remoto --docker "host=tcp://10.0.0.10:2376"
```

**Use quando:** quer apontar a CLI para outro daemon.

### `docker context inspect`

Mostra detalhes do context.

```bash
docker context inspect prod
```

**Use quando:** precisa validar endpoint e credenciais.

---

## 11) Segurança e hardening

### `docker scan`

Ferramenta histórica de análise de vulnerabilidades, dependente de disponibilidade/integração da plataforma.

```bash
docker scan minha-imagem:1.0.0
```

**Use quando:** quer checar vulnerabilidades, embora hoje muita gente prefira Trivy, Grype ou scanners de registry/CI.

### `--user`

Executar como não-root.

```bash
docker run --user 1000:1000 minha-app
```

**Use quando:** quer reduzir impacto em caso de exploração.

### `--cap-drop ALL`

Remove quase todos os privilégios.

```bash
docker run --cap-drop ALL minha-app
```

**Use quando:** quer um container mais seguro.

### `--security-opt no-new-privileges:true`

Bloqueia elevação de privilégios.

```bash
docker run --security-opt no-new-privileges:true minha-app
```

**Use quando:** quer reforçar o isolamento.

### `--read-only`

Filesystem somente leitura.

```bash
docker run --read-only minha-app
```

**Use quando:** quer impedir escrita inesperada.

### `--tmpfs`

Cria área temporária em memória.

```bash
docker run --tmpfs /tmp minha-app
```

**Use quando:** o app precisa escrever temporariamente sem persistir.

---

## 12) Depuração e troubleshooting

### `docker exec -it <container> sh`

Entra no container para inspeção.

```bash
docker exec -it web sh
```

**Use quando:** precisa checar arquivos, processos, DNS, rede ou logs internos.

### `docker logs -f`

Acompanha logs em tempo real.

```bash
docker logs -f web
```

**Use quando:** está rastreando inicialização ou erro intermitente.

### `docker inspect <container>`

Checa IP, mounts, rede, entrypoint, cmd e estado.

```bash
docker inspect web
```

**Use quando:** quer entender configuração real do container.

### `docker top <container>`

Mostra processos reais dentro do namespace do container.

```bash
docker top web
```

**Use quando:** quer validar PID 1, subprocessos e zumbis.

### `docker events`

Observa eventos do daemon.

```bash
docker events
```

**Use quando:** quer ver o que está acontecendo em tempo real no Docker.

### `docker stats`

Monitora CPU, memória e rede.

```bash
docker stats web
```

**Use quando:** suspeita de gargalo de recurso.

### `nsenter` + `docker inspect`

Não é comando Docker, mas é muito usado junto com Docker para entrar em namespaces do container.

**Use quando:** precisa investigar rede, processos ou mount namespace com precisão de nível kernel.

---

## 13) Boas práticas de uso

- Use **tags versionadas** em vez de `latest` em produção.
- Prefira **multi-stage build** para imagens menores.
- Use **non-root user** sempre que possível.
- Reduza privilégios com **capabilities mínimas**.
- Aplique **healthcheck** para detectar falhas cedo.
- Prefira **volumes** para dados persistentes.
- Use **bind mounts** em desenvolvimento.
- Limpe recursos com `prune` com cuidado para não apagar dados importantes.
- Use `docker inspect`, `logs`, `exec` e `stats` como trio principal de troubleshooting.

---

## 14) Mapa mental rápido: qual comando usar?

**Quero subir uma app localmente:** `docker run`, `docker compose up`

**Quero debugar um container:** `docker logs`, `docker exec`, `docker inspect`, `docker top`

**Quero mexer em imagens:** `docker build`, `docker pull`, `docker push`, `docker tag`, `docker rmi`

**Quero persistência:** `docker volume create`, `-v`, `--mount`

**Quero rede entre serviços:** `docker network create`, `connect`, `inspect`

**Quero ver consumo:** `docker stats`, `docker system df`

**Quero limpar o ambiente:** `docker system prune`, `docker container prune`, `docker image prune`

**Quero segurança:** `--user`, `--cap-drop`, `--security-opt`, `--read-only`

---

## 15) Comandos auxiliares úteis

### `docker login`

Autentica em um registry.

```bash
docker login
```

**Use quando:** vai puxar ou publicar imagens privadas.

### `docker logout`

Remove credenciais do registry.

```bash
docker logout
```

**Use quando:** quer sair de um registry.

### `docker manifest`

Trabalha com manifests multi-arquitetura.

```bash
docker manifest inspect nginx:latest
```

**Use quando:** quer entender suporte a multi-arch.

### `docker checkpoint`

Recursos avançados de checkpoint/restore, quando suportados.

```bash
docker checkpoint create web chk1
```

**Use quando:** trabalha com cenários experimentais ou avançados de runtime.

---

## 16) Resumo final por categoria

- **Imagem:** `build`, `pull`, `push`, `tag`, `history`, `save`, `load`
- **Container:** `run`, `ps`, `stop`, `start`, `restart`, `kill`, `exec`, `logs`, `inspect`
- **Rede:** `network ls/create/inspect/connect/disconnect/rm`
- **Volume:** `volume ls/create/inspect/rm/prune`
- **Limpeza:** `system df`, `system prune`, `image prune`, `container prune`, `builder prune`
- **Compose:** `compose up/down/ps/logs/exec/build/config`
- **Segurança:** `--user`, `--cap-drop`, `--security-opt`, `--read-only`
- **BuildKit:** `buildx`, `--secret`, `--ssh`, `--cache-from`, `--cache-to`, `--platform`

---

## 17) Observação importante

Este cheatsheet cobre os comandos mais usados e os mais importantes para estudo avançado, produção e troubleshooting. O Docker evolui com versões, então algumas flags e subcomandos podem variar levemente entre ambientes.

Se o objetivo for estudo para entrevista ou prática de senioridade, o melhor caminho é dominar primeiro:

1. `docker run`
2. `docker build`
3. `docker logs`
4. `docker exec`
5. `docker inspect`
6. `docker network`
7. `docker volume`
8. `docker compose`
9. `docker system prune`
10. `docker buildx`

---

Fim do cheatsheet.
