# 蛇梯游戏

### 功能性接口

- 生成随机棋盘

  - `mothed POST`

  - `url "http://127.0.0.1:8080/api/v1/randomboard"`

  - `request`

    ```json
    {
        "player_nums":4
    }
    ```

  - `response`

    - Good

      ```json
      {
          "code": 0,
          "msg": "Random Board Success",
          "body": {
              "Name": "ce7kba28ra55flrobnl0",
              "Position": [
                  [
                      100,
                      99,
                      98,
                      97,
                      96,
                      95,
                      94,
                      93,
                      92,
                      91
                  ],
                  [
                      81,
                      82,
                      83,
                      84,
                      85,
                      86,
                      87,
                      88,
                      89,
                      90
                  ],
                  [
                      80,
                      79,
                      78,
                      77,
                      76,
                      75,
                      74,
                      73,
                      72,
                      71
                  ],
                  [
                      61,
                      62,
                      63,
                      64,
                      65,
                      66,
                      67,
                      68,
                      69,
                      70
                  ],
                  [
                      60,
                      59,
                      58,
                      57,
                      56,
                      55,
                      54,
                      53,
                      52,
                      51
                  ],
                  [
                      41,
                      42,
                      43,
                      44,
                      45,
                      46,
                      47,
                      48,
                      49,
                      50
                  ],
                  [
                      40,
                      39,
                      38,
                      37,
                      36,
                      35,
                      34,
                      33,
                      32,
                      31
                  ],
                  [
                      21,
                      22,
                      23,
                      24,
                      25,
                      26,
                      27,
                      28,
                      29,
                      30
                  ],
                  [
                      20,
                      19,
                      18,
                      17,
                      16,
                      15,
                      14,
                      13,
                      12,
                      11
                  ],
                  [
                      1,
                      2,
                      3,
                      4,
                      5,
                      6,
                      7,
                      8,
                      9,
                      10
                  ]
              ],
              "Snake": [
                  [
                      43,
                      7
                  ],
                  [
                      14,
                      42
                  ],
                  [
                      62,
                      41
                  ],
                  [
                      42,
                      58
                  ],
                  [
                      14,
                      88
                  ],
                  [
                      54,
                      25
                  ],
                  [
                      94,
                      1
                  ],
                  [
                      57,
                      45
                  ],
                  [
                      72,
                      88
                  ],
                  [
                      64,
                      9
                  ]
              ],
              "Ladder": [
                  [
                      85,
                      52
                  ],
                  [
                      64,
                      67
                  ],
                  [
                      49,
                      36
                  ],
                  [
                      84,
                      41
                  ],
                  [
                      96,
                      0
                  ]
              ],
              "Player": [
                  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJib2FyZCI6ImNlN2tiYTI4cmE1NWZscm9ibmwwIiwiZXhwIjoxNjcwMzM3NDY0LCJwbGF5ZXIiOiJjZTdrYmEyOHJhNTVmbHJvYm5sZyJ9.p9v_pWVlAYTO18FMjO6LQNgC1St69715v_tMhA-dtV4",
                  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJib2FyZCI6ImNlN2tiYTI4cmE1NWZscm9ibmwwIiwiZXhwIjoxNjcwMzM3NDY0LCJwbGF5ZXIiOiJjZTdrYmEyOHJhNTVmbHJvYm5tMCJ9.AqLXEdGimCkHZHJYTJ_dtdTQsww6RaBqx5_5LkdEvks",
                  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJib2FyZCI6ImNlN2tiYTI4cmE1NWZscm9ibmwwIiwiZXhwIjoxNjcwMzM3NDY0LCJwbGF5ZXIiOiJjZTdrYmEyOHJhNTVmbHJvYm5tZyJ9.0TvnQ6u3wCwLc78cEKnaTKPE5nTPIVEIsbhjwCEsHwM",
                  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJib2FyZCI6ImNlN2tiYTI4cmE1NWZscm9ibmwwIiwiZXhwIjoxNjcwMzM3NDY0LCJwbGF5ZXIiOiJjZTdrYmEyOHJhNTVmbHJvYm5uMCJ9.1tlPezelHLRsyN57vJRLa6ZUkozi2mkQbSuij2MQXsc"
              ]
          }
      }
      ```

    - Bad

      ```json
      {"code":0,"msg":"Random Board Failed"}
      ```

- 随机投掷骰子

  - `method GET`

  - `url "http://127.0.0.1:8080/api/v1/admin/randomdice"`

  - 选择一个player 当作token 然后header 中传入 jwt 验证方式

  - `headers`

    ```json
    "Authorization":"Bear eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJib2FyZCI6ImNlN2tiYTI4cmE1NWZscm9ibmwwIiwiZXhwIjoxNjcwMzM3NDY0LCJwbGF5ZXIiOiJjZTdrYmEyOHJhNTVmbHJvYm5sZyJ9.p9v_pWVlAYTO18FMjO6LQNgC1St69715v_tMhA-dtV4"
    ```

  - `response`

    - Good

      ```json
      {
          "code": 0,
          "msg": "Random Dice Success",
          "body": 5
      }
      ```

    - Bad

      ```json
      {
          "code": 1,
          "msg": "Random Dice Bad",
          "body": 5
      }
      ```





- 存储走的路径