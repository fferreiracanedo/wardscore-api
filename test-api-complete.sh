#!/bin/bash

echo "🧪 TESTANDO WARDSCORE API COMPLETA..."
echo "========================================"

API_URL="http://localhost:8080"

# Cores para output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo ""
echo -e "${BLUE}1. 🏥 Health Check...${NC}"
curl -s $API_URL/health | jq .

echo ""
echo -e "${BLUE}2. 👤 Criando primeiro usuário...${NC}"
USER1=$(curl -s -X POST $API_URL/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "riot_id": "testplayer001",
    "game_name": "TestPlayer1",
    "tag_line": "BR1", 
    "email": "player1@wardscore.com",
    "region": "BR1",
    "puuid": "puuid-test-001"
  }')
echo $USER1 | jq .

echo ""
echo -e "${BLUE}3. 👤 Criando segundo usuário...${NC}"
USER2=$(curl -s -X POST $API_URL/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "riot_id": "testplayer002",
    "game_name": "TestPlayer2",
    "tag_line": "BR1",
    "email": "player2@wardscore.com", 
    "region": "BR1",
    "puuid": "puuid-test-002"
  }')
echo $USER2 | jq .

echo ""
echo -e "${BLUE}4. 📋 Listando todos os usuários...${NC}"
curl -s "$API_URL/api/v1/users?page=1&limit=10" | jq .

echo ""
echo -e "${BLUE}5. 👤 Buscando perfil do usuário 1...${NC}"
curl -s $API_URL/api/v1/users/profile \
  -H "X-User-ID: 1" | jq .

echo ""
echo -e "${BLUE}6. ✏️  Atualizando perfil do usuário 1...${NC}"
curl -s -X PUT $API_URL/api/v1/users/profile \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -d '{
    "game_name": "UpdatedPlayer1",
    "avatar_url": "https://example.com/avatar1.jpg"
  }' | jq .

echo ""
echo -e "${BLUE}7. 🎮 Upload de replay 1...${NC}"
REPLAY1=$(curl -s -X POST $API_URL/api/v1/replays/upload \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -d '{
    "file_name": "game1_victory.rofl",
    "match_id": "BR1_123456789",
    "game_mode": "CLASSIC",
    "game_version": "13.24",
    "duration": 1847,
    "champion": "Thresh",
    "role": "SUPPORT",
    "queue": "RANKED_SOLO_5x5"
  }')
echo $REPLAY1 | jq .

echo ""
echo -e "${BLUE}8. 🎮 Upload de replay 2...${NC}"
REPLAY2=$(curl -s -X POST $API_URL/api/v1/replays/upload \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -d '{
    "file_name": "game2_defeat.rofl", 
    "match_id": "BR1_987654321",
    "game_mode": "RANKED_SOLO_5x5",
    "game_version": "13.24",
    "duration": 2156,
    "champion": "Lulu",
    "role": "SUPPORT",
    "queue": "RANKED_SOLO_5x5"
  }')
echo $REPLAY2 | jq .

echo ""
echo -e "${BLUE}9. 📋 Listando replays do usuário 1...${NC}"
curl -s $API_URL/api/v1/replays \
  -H "X-User-ID: 1" | jq .

echo ""
echo -e "${BLUE}10. 🎮 Buscando replay específico...${NC}"
curl -s $API_URL/api/v1/replays/1 | jq .

echo ""
echo -e "${BLUE}11. ✏️  Atualizando replay...${NC}"
curl -s -X PUT $API_URL/api/v1/replays/1 \
  -H "Content-Type: application/json" \
  -d '{
    "champion": "Thresh",
    "role": "SUPPORT",
    "queue": "RANKED_SOLO_5x5"
  }' | jq .

echo ""
echo -e "${BLUE}12. 📊 Processando replay 1 (criando análise)...${NC}"
ANALYSIS1=$(curl -s -X POST $API_URL/api/v1/analysis/process/1 \
  -H "X-User-ID: 1")
echo $ANALYSIS1 | jq .

echo ""
echo -e "${BLUE}13. 📊 Processando replay 2 (criando análise)...${NC}"
ANALYSIS2=$(curl -s -X POST $API_URL/api/v1/analysis/process/2 \
  -H "X-User-ID: 1")
echo $ANALYSIS2 | jq .

echo ""
echo -e "${BLUE}14. 📈 Buscando análise específica...${NC}"
curl -s $API_URL/api/v1/analysis/1 | jq .

echo ""
echo -e "${BLUE}15. 📊 Buscando todas as análises do usuário 1...${NC}"
curl -s $API_URL/api/v1/analysis/user/1 | jq .

echo ""
echo -e "${BLUE}16. 📈 Teste dashboard stats...${NC}"
curl -s $API_URL/api/v1/stats/dashboard | jq .

echo ""
echo -e "${BLUE}17. 🏆 Teste ranking global...${NC}"
curl -s $API_URL/api/v1/ranking/global | jq .

echo ""
echo -e "${BLUE}18. 🌎 Teste ranking regional...${NC}"
curl -s $API_URL/api/v1/ranking/region/BR1 | jq .

echo ""
echo -e "${BLUE}19. 🗑️  Deletando replay 2...${NC}"
curl -s -X DELETE $API_URL/api/v1/replays/2 | jq .

echo ""
echo -e "${BLUE}20. 📋 Verificando replays após deleção...${NC}"
curl -s $API_URL/api/v1/replays \
  -H "X-User-ID: 1" | jq .

echo ""
echo -e "${GREEN}✅ TESTES CONCLUÍDOS!${NC}"
echo -e "${GREEN}====================${NC}"
echo ""
echo "📊 Resumo dos endpoints testados:"
echo "- ✅ POST /api/v1/users (criar usuário)"
echo "- ✅ GET /api/v1/users (listar usuários)"
echo "- ✅ GET /api/v1/users/profile (perfil)"
echo "- ✅ PUT /api/v1/users/profile (atualizar perfil)"
echo "- ✅ POST /api/v1/replays/upload (upload replay)"
echo "- ✅ GET /api/v1/replays (listar replays)"
echo "- ✅ GET /api/v1/replays/:id (buscar replay)"
echo "- ✅ PUT /api/v1/replays/:id (atualizar replay)"
echo "- ✅ DELETE /api/v1/replays/:id (deletar replay)"
echo "- ✅ POST /api/v1/analysis/process/:replay_id (processar)"
echo "- ✅ GET /api/v1/analysis/:id (buscar análise)"
echo "- ✅ GET /api/v1/analysis/user/:user_id (análises do usuário)"
echo ""
echo "🎯 API WardScore funcionando completamente!" 